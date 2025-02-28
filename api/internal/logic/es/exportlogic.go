package es

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/samber/lo"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/gql"
	"github.com/0x587/guardeye/common/gql/listener"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportLogic {
	return &ExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportLogic) Export(req *types.EsExportReq) (resp *types.EsExportRsp, err error) {
	tree, errs := gql.ParseQuery(req.Query)
	if len(errs) > 0 {
		return &types.EsExportRsp{
			QueryErrors: errs,
		}, nil
	}
	s := gql.ScheduleQuery(tree)
	filter, err := makeFilter(s)
	if err != nil {
		return nil, err
	}
	colNames := lo.Map(s.Result, func(r *listener.ResultEntry, _ int) string { return r.Alias })

	taskId := uuid.New().String()
	redisKey := fmt.Sprintf("report_task_%s", taskId)
	if err := l.svcCtx.Redis.HsetCtx(l.ctx, redisKey, "status", "running"); err != nil {
		return nil, err
	}
	if err := l.svcCtx.Redis.HsetCtx(l.ctx, redisKey, "process", "0"); err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)

	go func() {
		// 异步拉取es数据写入minio
		err := func() error {
			resCh, err := fetchEs(ctx, l.svcCtx.Es, filter)
			if err != nil {
				return err
			}
			// 生成临时文件
			tempFile, err := os.CreateTemp("", "export_temp_")
			if err != nil {
				return err
			}
			defer func() {
				_ = tempFile.Close()
				_ = os.Remove(tempFile.Name())
			}()
			var w exportWriter
			w = &csvImpl{out: tempFile}
			if err := w.WriteMeta(colNames); err != nil {
				return err
			}

			c := 0
			for res := range resCh {
				for _, r := range res.records {
					c += 1
					ij := gql.NewInjector(r.Fields.Message[0])
					rs := lo.FilterMap(s.Result, func(result *listener.ResultEntry, _ int) (string, bool) {
						v, err := result.Value.Vf(ij)
						return fmt.Sprintf("%v", v), err == nil
					})
					if err := w.Write(rs); err != nil {
						return err
					}
				}

				if err := l.svcCtx.Redis.HsetCtx(ctx, redisKey, "total", strconv.Itoa(res.total.Value)); err != nil {
					logx.Error(err)
				}
				if err := l.svcCtx.Redis.HsetCtx(ctx, redisKey, "status", "running"); err != nil {
					logx.Error(err)
				}
				if err := l.svcCtx.Redis.HsetCtx(ctx, redisKey, "process", strconv.Itoa(c)); err != nil {
					logx.Error(err)
				}
			}
			_, err = l.svcCtx.Minio.FPutObject(ctx, "export", fmt.Sprintf("/%s.csv", taskId), tempFile.Name(), minio.PutObjectOptions{})
			if err != nil {
				return err
			}
			if err := l.svcCtx.Redis.HsetCtx(ctx, redisKey, "done", "1"); err != nil {
				logx.Error(err)
			}
			return nil
		}()

		if err != nil {
			if e := l.svcCtx.Redis.HsetCtx(ctx, redisKey, "err", err.Error()); e != nil {
				logx.Error(err)
			}
			if err := l.svcCtx.Redis.HsetCtx(ctx, redisKey, "status", "fail"); err != nil {
				logx.Error(err)
			}
		}
	}()

	return &types.EsExportRsp{
		QueryErrors: errs,
		TaskId:      taskId,
	}, nil
}

type exportWriter interface {
	WriteMeta(colNames []string) error
	Write(row []string) error
}
type csvImpl struct {
	out io.StringWriter
}

func (i *csvImpl) WriteMeta(colNames []string) error {
	_, err := i.out.WriteString(strings.Join(colNames, ",") + "\n")
	return err
}

func (i *csvImpl) Write(row []string) error {
	_, err := i.out.WriteString(strings.Join(row, ",") + "\n")
	return err
}

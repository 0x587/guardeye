package parselogkey

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
	"github.com/zeromicro/go-queue/kq"
	"gopkg.in/yaml.v3"
)

func New(ctx context.Context, svcCtx *svc.ServiceContext) kq.ConsumeHandler {
	return &impl{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type impl struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (i *impl) Consume(ctx context.Context, key, val string) error {
	d := &report.MQLog{}
	if err := json.Unmarshal([]byte(val), d); err != nil {
		return err
	}
	var keys []string
	switch d.GetLog().GetType() {
	case report.LogType_TEXT:
		keys = []string{"TEXT"}
	case report.LogType_YAML:
		keys = parseYaml(d.GetLog().GetMessage())
	}
	return i.svcCtx.DataKeyRedisClient.SetKey(
		ctx,
		d.GetNodeInfo(),
		d.GetLog().GetProvider(),
		keys,
	)
}

func parseYaml(msg string) []string {
	var data interface{}
	err := yaml.Unmarshal([]byte(msg), &data)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// 用于存储所有路径
	var paths []string
	// 从根开始递归提取路径
	extractPaths(data, "", &paths)
	return paths
}

func extractPaths(value interface{}, prefix string, paths *[]string) {
	// 获取值的类型
	val := reflect.ValueOf(value)

	// 如果值是一个map，则递归遍历
	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			// 获取键对应的值
			fieldVal := val.MapIndex(key)
			// 构建路径
			newPrefix := fmt.Sprintf("%s.%v", prefix, key)
			// 递归处理
			extractPaths(fieldVal.Interface(), newPrefix, paths)
		}
	case reflect.Struct:
		// 如果是结构体类型，获取字段信息并递归处理
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			// 获取字段值
			fieldVal := val.Field(i)
			// 跳过未导出字段
			if field.PkgPath != "" {
				continue
			}
			// 递归处理
			newPrefix := fmt.Sprintf("%s.%s", prefix, field.Name)
			extractPaths(fieldVal.Interface(), newPrefix, paths)
		}
	case reflect.Slice:
		// 如果是切片，遍历每个元素
		for i := 0; i < val.Len(); i++ {
			// 获取切片中的每个元素的值
			elementVal := val.Index(i)
			// 构建新的路径
			newPrefix := fmt.Sprintf("%s[%d]", prefix, i)
			// 递归处理
			extractPaths(elementVal.Interface(), newPrefix, paths)
		}
	default:
		// 如果值是一个基本类型并且不是空值，则添加路径
		if val.IsValid() && !isEmpty(val) {
			*paths = append(*paths, prefix[1:]) // 去除最前面的点
		}
	}
}

// 判断值是否为空
func isEmpty(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Map, reflect.Array, reflect.Slice:
		return val.Len() == 0
	case reflect.Ptr:
		return val.IsNil()
	default:
		return false
	}
}

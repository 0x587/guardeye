package svcctx

import (
	"github.com/0x587/guardeye/common/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	NodeDBClient        model.NodeModel
	RawLogDBClient      model.RawlogModel
	LogToMetricDBClient model.LogToMetricQueryModel
}

func NewServiceContext(dbConn sqlx.SqlConn) *ServiceContext {
	return &ServiceContext{
		NodeDBClient:        model.NewNodeModel(dbConn),
		RawLogDBClient:      model.NewRawlogModel(dbConn),
		LogToMetricDBClient: model.NewLogToMetricQueryModel(dbConn),
	}
}

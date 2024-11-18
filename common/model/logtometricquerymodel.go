package model

import (
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
)

var _ LogToMetricQueryModel = (*customLogToMetricQueryModel)(nil)

type (
	// LogToMetricQueryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLogToMetricQueryModel.
	LogToMetricQueryModel interface {
		logToMetricQueryModel
		withSession(session sqlx.Session) LogToMetricQueryModel
		FindForLog(ctx context.Context, clientId, provider string) ([]*LogToMetricQuery, error)
	}

	customLogToMetricQueryModel struct {
		*defaultLogToMetricQueryModel
	}
)

func (m *customLogToMetricQueryModel) FindForLog(ctx context.Context, clientId, provider string) ([]*LogToMetricQuery, error) {
	query := fmt.Sprintf("select %s from %s where client_id = $1 and provider = $2", logToMetricQueryRows, m.table)
	var resp []*LogToMetricQuery
	err := m.conn.QueryRowsCtx(ctx, &resp, query, clientId, provider)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewLogToMetricQueryModel returns a model for the database table.
func NewLogToMetricQueryModel(conn sqlx.SqlConn) LogToMetricQueryModel {
	return &customLogToMetricQueryModel{
		defaultLogToMetricQueryModel: newLogToMetricQueryModel(conn),
	}
}

func (m *customLogToMetricQueryModel) withSession(session sqlx.Session) LogToMetricQueryModel {
	return NewLogToMetricQueryModel(sqlx.NewSqlConnFromSession(session))
}

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
		List(ctx context.Context) ([]*LogToMetricQuery, error)
		ListForNode(ctx context.Context, clientId string) ([]*LogToMetricQuery, error)
		ListForLog(ctx context.Context, clientId, provider string) ([]*LogToMetricQuery, error)
	}

	customLogToMetricQueryModel struct {
		*defaultLogToMetricQueryModel
	}
)

func (m *customLogToMetricQueryModel) List(ctx context.Context) ([]*LogToMetricQuery, error) {
	query := fmt.Sprintf("select %s from %s", logToMetricQueryRows, m.table)
	return m.list(ctx, query)
}

func (m *customLogToMetricQueryModel) ListForNode(ctx context.Context, clientId string) ([]*LogToMetricQuery, error) {
	query := fmt.Sprintf("select %s from %s where client_id = $1", logToMetricQueryRows, m.table)
	return m.list(ctx, query, clientId)
}

func (m *customLogToMetricQueryModel) ListForLog(ctx context.Context, clientId, provider string) ([]*LogToMetricQuery, error) {
	query := fmt.Sprintf("select %s from %s where client_id = $1 and provider = $2", logToMetricQueryRows, m.table)
	return m.list(ctx, query, clientId, provider)
}

func (m *customLogToMetricQueryModel) list(ctx context.Context, query string, args ...interface{}) ([]*LogToMetricQuery, error) {
	var resp []*LogToMetricQuery
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
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

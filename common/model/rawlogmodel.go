package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RawlogModel = (*customRawlogModel)(nil)

type (
	// RawlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRawlogModel.
	RawlogModel interface {
		rawlogModel
		withSession(session sqlx.Session) RawlogModel
		ListLastSeen(ctx context.Context) ([]*Rawlog, error)
	}

	customRawlogModel struct {
		*defaultRawlogModel
	}
)

func (m *customRawlogModel) ListLastSeen(ctx context.Context) ([]*Rawlog, error) {
	query := fmt.Sprintf(`SELECT DISTINCT ON (client_id) * from %s order by client_id, created_at desc`, m.table)
	var resp []*Rawlog
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch {
	case err == nil:
		return resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewRawlogModel returns a model for the database table.
func NewRawlogModel(conn sqlx.SqlConn) RawlogModel {
	return &customRawlogModel{
		defaultRawlogModel: newRawlogModel(conn),
	}
}

func (m *customRawlogModel) withSession(session sqlx.Session) RawlogModel {
	return NewRawlogModel(sqlx.NewSqlConnFromSession(session))
}

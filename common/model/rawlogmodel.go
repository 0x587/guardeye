package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RawlogModel = (*customRawlogModel)(nil)

type (
	// RawlogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRawlogModel.
	RawlogModel interface {
		rawlogModel
		withSession(session sqlx.Session) RawlogModel
		GetLastSeen(ctx context.Context, clientId uuid.UUID) (*Rawlog, error)
	}

	customRawlogModel struct {
		*defaultRawlogModel
	}
)

func (m *customRawlogModel) GetLastSeen(ctx context.Context, clientId uuid.UUID) (*Rawlog, error) {
	query := fmt.Sprintf(`SELECT * from %s where client_id = $1 order by created_at desc limit 1`, m.table)
	var resp Rawlog
	err := m.conn.QueryRowCtx(ctx, &resp, query, clientId.String())
	switch {
	case err == nil:
		return &resp, nil
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

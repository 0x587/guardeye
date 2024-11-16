package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NodeModel = (*customNodeModel)(nil)

type (
	// NodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNodeModel.
	NodeModel interface {
		nodeModel
		withSession(session sqlx.Session) NodeModel
		FindOneWithClientID(ctx context.Context, cid uuid.UUID) (*Node, error)
		ListGroupByClientID(ctx context.Context) ([]*Node, error)
		FindOneLast(ctx context.Context, cid uuid.UUID) (*Node, error)
	}

	customNodeModel struct {
		*defaultNodeModel
	}
)

func (m *customNodeModel) FindOneLast(ctx context.Context, cid uuid.UUID) (*Node, error) {
	query := fmt.Sprintf("select %s from %s where client_id = $1 order by created_at desc limit 1", nodeRows, m.table)
	var resp Node
	err := m.conn.QueryRowCtx(ctx, &resp, query, cid)
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customNodeModel) ListGroupByClientID(ctx context.Context) ([]*Node, error) {
	query := fmt.Sprintf(`SELECT DISTINCT ON (client_id) * FROM %s ORDER BY client_id, created_at DESC`, m.table)
	var resp []*Node
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

func (m *customNodeModel) FindOneWithClientID(ctx context.Context, cid uuid.UUID) (*Node, error) {
	query := fmt.Sprintf("select %s from %s where client_id = $1 order by created_at limit 1", nodeRows, m.table)
	var resp Node
	err := m.conn.QueryRowCtx(ctx, &resp, query, cid)
	switch {
	case err == nil:
		return &resp, nil
	case errors.Is(err, sqlx.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewNodeModel returns a model for the database table.
func NewNodeModel(conn sqlx.SqlConn) NodeModel {
	return &customNodeModel{
		defaultNodeModel: newNodeModel(conn),
	}
}

func (m *customNodeModel) withSession(session sqlx.Session) NodeModel {
	return NewNodeModel(sqlx.NewSqlConnFromSession(session))
}

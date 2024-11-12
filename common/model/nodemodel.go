package model

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ NodeModel = (*customNodeModel)(nil)

type (
	// NodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNodeModel.
	NodeModel interface {
		nodeModel
		FindOneWithClientID(ctx context.Context, clientID string) (*Node, error)
	}

	customNodeModel struct {
		*defaultNodeModel
	}
)

func (c *customNodeModel) FindOneWithClientID(ctx context.Context, clientID string) (*Node, error) {
	var data Node
	err := c.conn.FindOne(ctx, &data, bson.M{"clientID": clientID}, &options.FindOneOptions{
		Sort: bson.M{"createAt": -1},
	})
	switch {
	case err == nil:
		return &data, nil
	case errors.Is(err, mon.ErrNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewNodeModel returns a model for the mongo.
func NewNodeModel(url, db, collection string) NodeModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customNodeModel{
		defaultNodeModel: newDefaultNodeModel(conn),
	}
}

package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ NodeModel = (*customNodeModel)(nil)

type (
	// NodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNodeModel.
	NodeModel interface {
		nodeModel
	}

	customNodeModel struct {
		*defaultNodeModel
	}
)

// NewNodeModel returns a model for the mongo.
func NewNodeModel(url, db, collection string) NodeModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customNodeModel{
		defaultNodeModel: newDefaultNodeModel(conn),
	}
}

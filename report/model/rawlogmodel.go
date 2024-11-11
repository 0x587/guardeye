package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ RawLogModel = (*customRawLogModel)(nil)

type (
	// RawLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRawLogModel.
	RawLogModel interface {
		rawLogModel
	}

	customRawLogModel struct {
		*defaultRawLogModel
	}
)

// NewRawLogModel returns a model for the mongo.
func NewRawLogModel(url, db, collection string) RawLogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customRawLogModel{
		defaultRawLogModel: newDefaultRawLogModel(conn),
	}
}

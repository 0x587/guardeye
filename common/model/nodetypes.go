package model

import (
	"time"

	"github.com/0x587/guardeye/report/report"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Node struct {
	ID          primitive.ObjectID      `bson:"_id,omitempty" json:"id,omitempty"`
	ClientID    string                  `bson:"clientID,omitempty" json:"clientID,omitempty"`
	Description *report.NodeDescription `bson:"description,omitempty" json:"description,omitempty"`
	Ips         []string                `bson:"ip,omitempty" json:"ip,omitempty"`
	UpdateAt    time.Time               `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time               `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

package model

import (
	"time"

	"github.com/0x587/guardeye/report/report"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RawLog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ClientID string             `bson:"clientID,omitempty" json:"clientID,omitempty"`
	Message  string             `bson:"message,omitempty" json:"message,omitempty"`
	Provider *report.Provider   `bson:"provider,omitempty" json:"provider,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

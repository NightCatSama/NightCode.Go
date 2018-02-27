package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Account    string        `bson:",omitempty" json:"account"`
	Password   string        `bson:",omitempty" json:"-"`
	Email      string        `bson:",omitempty" json:"email,omitempty"`
	CreateTime time.Time     `bson:"create_time" json:"create_time"`
	UpdateTime time.Time     `bson:"update_time,omitempty" json:"update_time,omitempty"`
}

package models


import (
	"gopkg.in/mgo.v2/bson"
)


type Profile struct {
	ID        		bson.ObjectId `bson:"_id,omitempty"`
	UserID			int64 `bson:"user_id,omitempty"`
	CityID			int64 `bson:"city_id,omitempty"`
}
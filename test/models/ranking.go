package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Ranking struct{
	Id				bson.ObjectId  	`bson:"_id,omitempty"`
	City_id			int 			`bson:"city_id,omitempty"`
	Citizne_count	int				`bson:"citizen_count,omitempty"`
	Complaint_count	int				`bson:"complaint_count,omitempty"`
}
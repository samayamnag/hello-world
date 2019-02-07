package repositories

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"tutorial/test/models"
)


type ProfileRepository struct {}

type Profiles []models.Profile

func (p ProfileRepository) GetProfiles() []models.Profile {
	session, err := mgo.Dial("localhost")

	if err != nil {
	 	fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB("swachh_manch").C("profiles")
	results := Profiles{}

	if err := c.Find(nil).Limit(4).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
  }

  return results
}

func (p ProfileRepository) GroupByCity() []bson.M {
	session, err := mgo.Dial("localhost")

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo Server:", err)
	}

	defer session.Close()

	c := session.DB("swachh_manch").C("profiles")
	
	pipeline := []bson.M{
		bson.M{
			"$group": bson.M{
				"_id": "$city_id",
				"citizen_count": bson.M{ "$sum": 1 },
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
				"city_id": "$_id",
				"citizen_count": 1,
			},
		},
	}

	pipe := c.Pipe(pipeline)
	results := []bson.M{}
	err = pipe.All(&results)
	if err != nil {
		fmt.Println("Failed to write results: ", err)
	}
	return results
}
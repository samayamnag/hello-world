package repositories

import (
	"fmt"
	"tutorial/test/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"tutorial/test/models"
)

type RankingRepository struct{}


func (r RankingRepository) Insert(data []bson.M)  {
	db := database.DbConnect("swachh_bharat")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	stmIns, err := db.Prepare("INSERT INTO sbm_ranking(city_id, citizen_count) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmIns.Close()

	totalRecords := len(data)

	for i := 0; i < totalRecords; i++ {
		_, err := stmIns.Exec(data[i]["city_id"], data[i]["citizen_count"])
		
		if err != nil{
			panic(err.Error())
		}
	}
}

func (r RankingRepository) InsertToMongo(data []bson.M) {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err.Error())
	}

	defer session.Close()

	totalRecords := len(data)
	
	c := session.DB("swachh_manch").C("ranking")

	for i := 0; i < totalRecords; i++ {
		c.Insert(data[i])
	}
}

func (r RankingRepository) UpdateComplaintCount(data []GroupByCity) {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err.Error())
	}

	defer session.Close()
	ranking := models.Ranking{}

	c := session.DB("swachh_manch").C("ranking")

	for _, row := range data {
		err = c.Find(bson.M{"city_id": row.City_id}).One(&ranking)

		if err != nil {
			if(err.Error() == "not found") {
				fmt.Println("Not found", row.City_id)
				session.DB("swachh_manch").C("ranking").Insert(&models.Ranking{City_id: row.City_id, Complaint_count: row.Complaint_count})				
			}
		} else {	
			session.DB("swachh_manch").C("ranking").Update(bson.M{"city_id": row.City_id}, bson.M{"$set": bson.M{"complaint_count": row.Complaint_count}})
			fmt.Println("Found", row.City_id, ranking)
		}
	}
}
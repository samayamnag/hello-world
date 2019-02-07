package repositories

import (
	_ "fmt"
	"tutorial/test/models"
	"tutorial/test/database"
)


type ComplaintRepository struct{}

type Complaints []models.Complaint

type GroupByCity struct{
	City_id			int
	Complaint_count	int
}

func (c ComplaintRepository) Index() Complaints{
	db := database.DbConnect("swachh_bharat")
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}
	
	rows, err := db.Query("SELECT id, user_id, city_id, created_at FROM sbm_complaints")

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	complaint := models.Complaint{}
	complaints := []models.Complaint{}

	for rows.Next() {
		var comp models.Complaint
		err := rows.Scan(&comp.Id, &comp.User_id, &comp.City_id, &comp.Created_at)

		if err != nil {
			panic(err.Error())
		}

		complaint.Id = comp.Id
		complaint.User_id = comp.User_id
		complaint.City_id = comp.City_id
		complaint.Created_at = comp.Created_at

		complaints = append(complaints, complaint)
	}

	return complaints
}

func (c ComplaintRepository) GroupByCity() []GroupByCity{
	db := database.DbConnect("swachh_bharat")

	defer db.Close()
	err := db.Ping()

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT city_id, COUNT(*) as complaint_count FROM sbm_complaints GROUP BY city_id")

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	comlaintGroupBy := GroupByCity{}
	arrComplaintGroupBy := []GroupByCity{}

	for rows.Next() {
		var cityId int
		var count int
		err := rows.Scan(&cityId, &count)

		if err != nil {
			panic(err.Error())
		}

		comlaintGroupBy.City_id = cityId
		comlaintGroupBy.Complaint_count = count

		arrComplaintGroupBy = append(arrComplaintGroupBy, comlaintGroupBy)
		
	}
	return arrComplaintGroupBy
}
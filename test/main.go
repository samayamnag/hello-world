package main

import (
	"fmt"

	_ "tutorial/test/models"
	"tutorial/test/repositories"
	"tutorial/test/handlers"

	"github.com/spf13/viper"

)

var controller = &handlers.ProfileHandler{Repository: repositories.ProfileRepository{}}
var profileRepo = repositories.ProfileRepository{}
var rankingRepo = repositories.RankingRepository{}
var complaintRepo = repositories.ComplaintRepository{}

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString(`app.env`))
	fmt.Println(viper.GetString("app.env"))
	fmt.Println(viper.GetBool("app.debug"))
}

func Add(value1 int, value2 int) int {
    return value1 + value2
}

func Subtract(value1 int, value2 int) int {
    return value1 - value2
}

func main() { 
	/*data := profileRepo.GetProfiles()

	for _, v := range data {		
			fmt.Println(v.ID.Hex())
			fmt.Println(v.UserID)	
	}
	*/

	group := profileRepo.GroupByCity()
	fmt.Println(group)

	for _, v := range group {	
		fmt.Printf("City ID:%d Profile Count: %d\n", v["city_id"], v["citizen_count"])
	}	

	//rankingRepo.Insert(group)
	//rankingRepo.InsertToMongo(group)

	/*complaints := complaintRepo.Index()
	fmt.Println("=== Printing complaints ===")

	for _, complaint := range complaints {
		fmt.Println(complaint)
	} */

	cityComplaints := complaintRepo.GroupByCity()

	/*fmt.Println("=== Print City complaints ===")

	for _, v := range cityComplaints{
		fmt.Println(v)
	} */

	rankingRepo.UpdateComplaintCount(cityComplaints)

}
package handlers

import (
	"fmt"
	"tutorial/test/repositories"
	"tutorial/test/models"
	
)

type ProfileHandler struct {
	Repository repositories.ProfileRepository
}

type Profiles []models.Profile

func (c *ProfileHandler) Index(){
	profiles := c.Repository.GetProfiles()

	for _, profile := range profiles {
		fmt.Println(profile.ID)
		fmt.Println(profile.UserID)
	}
}
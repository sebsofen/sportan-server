package handlers

import (

	"sportan/databases"
	"sportan/repositories"
	"sportan/services"

	"fmt"
)

type SportHandler struct {
	userR   *repositories.UserRepository
	imageR *repositories.ImageRepository
	repo *repositories.SportRepository
	metric *databases.MetricApi
}

func NewSportHandler(repo *repositories.SportRepository,userR *repositories.UserRepository, imageR *repositories.ImageRepository, metricApi *databases.MetricApi) *SportHandler {
	return &SportHandler{
		repo:   repo,
		userR : userR,
		imageR : imageR,
		metric: metricApi,
	}
}

//Create and store user in database
//string is provided by device, username by server
func (ch *SportHandler) CreateSport(token string, sport *services.Sport) (*services.Sport, error) {
	userid, _ := ch.userR.GetUserIdFromToken(token)

	if userid != "" {
		if sport.Icon != nil && sport.Icon.Content != nil {
			//image was transmitted with sport.
			//create new image
			ch.repo.CreateSport(sport)
		}
	}

	return sport,nil

}

func (ch *SportHandler) GetAllSports(bla string) ([]*services.Sport, error) {
	sports := ch.repo.GetAllSports()

	for _, sport := range sports {
		fmt.Println(*sport.Name)

	}
	return sports,nil
}

func (ch *SportHandler) GetSportById(token string, sportid string) (*services.Sport, error) {
	_, err := ch.userR.GetUserIdFromToken(token)
	if err != nil {
		return nil, err
	}

	sport := ch.repo.GetSportById(sportid)

	return sport, nil
}
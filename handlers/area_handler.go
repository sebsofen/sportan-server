package handlers

import (
	"sportan/databases"
	"sportan/repositories"
	"sportan/services"

	"github.com/nu7hatch/gouuid"
	"fmt"
)

type AreaHandler struct {
	repo   *repositories.AreaRepository
	metric *databases.MetricApi
	userR *repositories.UserRepository
}

func NewAreaHandler(mRepo *repositories.AreaRepository, userR *repositories.UserRepository, metricApi *databases.MetricApi) *AreaHandler {
	return &AreaHandler{
		repo:   mRepo,
		metric: metricApi,
		userR : userR,
	}
}



func (ch *AreaHandler) CreateArea(token string, area *services.Area) error {
	userid, err := ch.userR.GetUserIdFromToken(token)
	if err != nil {
		fmt.Println("Create Area - no user")
		return err
	}
	if ch.userR.IsAdmin(userid) {
		u, _ := uuid.NewV4()
		id := u.String()
		if area.Center == nil {
			centerCoord := []float64{0.0, 0.0}

			for _, v := range area.Coords {
				centerCoord[1] += v.Lat
				centerCoord[0] += v.Lon
			}
			centerCoord[0] /= float64(len(area.Coords))
			centerCoord[1] /= float64(len(area.Coords))
			area.Center = centerCoord
		}
		area.ID = &id



		return ch.repo.CreateArea(area)

	}else {
		return  nil
	}


}

func (ch *AreaHandler) GetAllAreasInCity(areaid string) ([]string, error) {
	areas, err := ch.repo.GetAllAreasInCity(areaid)

	if err != nil {
		return nil, err;
	}

	sAreas := make([]string, len(areas))

	for i, area := range areas {
		sAreas[i] = *area.ID
	}
	return sAreas, nil
}

func (ch *AreaHandler) GetAreaById(token string,id string) (*services.Area, error) {
	/*_, err := ch.userR.GetUserIdFromToken(token)
	if err != nil{
		return nil, err
	}*/
	return ch.repo.GetAreaById(id), nil
}


func (ch *AreaHandler) UpdateArea(token string, area *services.Area) error {
	userid, err := ch.userR.GetUserIdFromToken(token)
	if err != nil {
		return err
	}
	if ch.userR.IsAdmin(userid) {
		if(area.ID == nil) {

			return ch.CreateArea(token,area)
		}else{
			return ch.repo.UpdateArea(area)
		}


	}

	return nil
}

func (ch *AreaHandler) DeleteArea(token string, area *services.Area) error {
	userid, _ := ch.userR.GetUserIdFromToken(token)

	if ch.userR.IsAdmin(userid) {
		ch.repo.UpdateArea(area)

	}

	return nil
}



func (ch *AreaHandler) GetNearBy(token string, coord *services.Coordinate, limit int32) ([]string, error) {

	if limit > 500 {
		limit = 500
	}

	areas, _ := ch.repo.GetNearBy(coord.Lon, coord.Lat, limit)
	sAreas := make([]string, len(areas))

	for i, area := range areas {
		sAreas[i] = *area.ID

	}
	return sAreas, nil

}

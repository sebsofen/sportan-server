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
	userid, _ := ch.userR.GetUserIdFromToken(token)

	if ch.userR.IsAdmin(userid) {
		u, _ := uuid.NewV4()
		id := u.String()
		centerCoord := &services.Coordinate{
			Lat: 0.0,
			Lon: 0.0,
		}

		for _, v := range area.Coords {
			centerCoord.Lat += v.Lat
			centerCoord.Lon += v.Lon
		}
		centerCoord.Lat /= float64(len(area.Coords))
		centerCoord.Lon /= float64(len(area.Coords))
		area.Center = centerCoord
		area.ID = &id

		ch.repo.CreateArea(area)

		return nil

	}else {
		return  nil
	}


}

func (ch *AreaHandler) GetAllAreasInCity(areaid string) ([]*services.Area, error) {
	return ch.repo.GetAllAreasInCity(areaid)
}

func (ch *AreaHandler) GetAreaById(id string) (*services.Area, error) {
	return ch.repo.GetAreaById(id), nil
}


func (ch *AreaHandler) UpdateArea(token string, area *services.Area) error {
	fmt.Println("tolken " + token)
	userid, _ := ch.userR.GetUserIdFromToken(token)

	if ch.userR.IsAdmin(userid) {
		ch.repo.UpdateArea(area)

	}

	return nil
}



func (ch *AreaHandler) GetNearBy(coord *services.Coordinate, limit int32) ([]*services.Area, error) {

	if limit > 100 {
		limit = 100
	}

	areas, _ := ch.repo.GetNearBy(coord.Lon, coord.Lat, limit)
	sAreas := make([]*services.Area, len(*areas))

	for i, area := range *areas {
		sAreas[i] = &services.Area{
			Title: &area.Title,
			ID:    &area.ID,
		}

	}
	return sAreas, nil

}

package handlers

import (
	"sportan/databases"
	"sportan/repositories"
	"sportan/services"

	"github.com/nu7hatch/gouuid"
)

type AreaHandler struct {
	repo   *repositories.AreaRepository
	metric *databases.MetricApi
}

func NewAreaHandler(mRepo *repositories.AreaRepository, metricApi *databases.MetricApi) *AreaHandler {
	return &AreaHandler{
		repo:   mRepo,
		metric: metricApi,
	}
}

//Create and store user in database
//string is provided by device, username by server
func (ch *AreaHandler) CreateArea(title string, description string, coordslist []*services.Coordinate) error {
	//create userid
	u, _ := uuid.NewV4()
	id := u.String()
	centerCoord := &services.Coordinate{
		Lat: 0.0,
		Lon: 0.0,
	}

	for _, v := range coordslist {
		centerCoord.Lat += v.Lat
		centerCoord.Lon += v.Lon
	}
	centerCoord.Lat /= float64(len(coordslist))
	centerCoord.Lon /= float64(len(coordslist))
	err := ch.repo.CreateArea(id, title, description, coordslist, centerCoord)
	if err != nil {
		panic(err)
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
			Title: area.Title,
			ID:    area.ID,
		}

	}
	return sAreas, nil

}

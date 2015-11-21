package handlers

import (
	"sportan/databases"
	"sportan/repositories"
	"sportan/services"
"github.com/nu7hatch/gouuid"
)

type CityHandler struct {
	repo   *repositories.CityRepository
	metric *databases.MetricApi
}

func NewCityHandler(mRepo *repositories.CityRepository, metricApi *databases.MetricApi) *CityHandler {
	return &CityHandler{
		repo:   mRepo,
		metric: metricApi,
	}
}

//Create and store user in database
//string is provided by device, username by server
func (ch *CityHandler) CreateCity(token string,title string, coordslist []*services.Coordinate) error {
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

	err := ch.repo.CreateCity(id, title, coordslist, centerCoord)
	if err != nil {
		panic(err)
	}
	return nil
}

func (ch *CityHandler) GetCityById(id string) (*services.City,error) {
	return ch.repo.GetCityById(id),nil
}

func (ch *CityHandler) GetNearBy(coord *services.Coordinate, limit int32) ([]*services.City, error) {
	return ch.repo.GetNearBy(coord.Lat,coord.Lon, limit)
}

func (ch *CityHandler) GetAllCities() ([]*services.City,error) {
	return ch.repo.GetAllCities()
}



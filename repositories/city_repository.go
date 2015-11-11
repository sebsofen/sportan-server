package repositories

import (
	"fmt"
	"sportan/databases"
	"sportan/services"
	"gopkg.in/mgo.v2/bson"
)

type City struct {
	Name       string                 `bson:"name,omitempty"`
	Coords      []*services.Coordinate `bson:"coords,omitempty"`
	Center      []float64              `bson:"center,omitempty"`
	ID          string                 `bson:"cityid,omitempty"`
}

type CityRepository struct {
	mongo *databases.MongoConfig
}

func NewCityRepository(mConfig *databases.MongoConfig) *CityRepository {
	//initialize the area, if not exists ;-)
	//Collection is name

	return &CityRepository{
		mongo: mConfig,
	}
}

func (rep *CityRepository) CreateCity(id string, name string, coordslist []*services.Coordinate, center *services.Coordinate) error {
	fmt.Println(name)
	//save center in database for faster access
	err := rep.mongo.Collection.Insert(
		&City{
			ID:          id,
			Name:        name,
			Center:      []float64{center.Lon, center.Lat},
			Coords:      coordslist})

	return err
}

func (rep *CityRepository) GetCityById(id string) *services.City {
	var city City
	err := rep.mongo.Collection.Find(bson.M{"cityid": id}).One(&city)

	if(err != nil){
		panic(err)
	}

	sCoords := make([]*services.Coordinate, len(city.Coords))

	for i, coord := range city.Coords {
		fmt.Println(i)
		sCoords[i] = &services.Coordinate{
			Lat: coord.Lat,
			Lon: coord.Lon,
		}

	}

	return &services.City{
		ID: id,
		Coords: sCoords,
		Center: &services.Coordinate{
			Lat: 0,
			Lon: 1,
			},
		Name: city.Name,
	}
}

func (rep *CityRepository) GetNearBy(longitude float64, latitude float64, limit int32) ([]*services.City, error) {
	fmt.Println("entry - get city by id")
	var cities []City
	err := rep.mongo.Collection.Find(
		bson.M{"center":
					bson.M{"$near":
						bson.M{"$geometry":
							bson.M{"type":        "Point",
										"coordinates": []float64{longitude, latitude}}}}}).Limit(int(limit)).All(&cities)
	if err != nil {
		panic(err)
	}

	//city service.city converter!
	cCities := make([]*services.City, len(cities))
	for i, city := range cities {
		cCities[i] = &services.City{
			ID : city.ID,
			Coords: city.Coords,
			Center: &services.Coordinate{
				Lat: city.Center[0],
				Lon: city.Center[1],
			},
			Name : city.Name,
		}
	}
	fmt.Println("exit - get city by id")
	return cCities, nil
}

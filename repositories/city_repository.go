package repositories

import (
	"fmt"
	"sportan/databases"
	"sportan/services"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
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

	err :=mConfig.Collection.EnsureIndex(
		mgo.Index{
			Key: []string{"$2dsphere:center"}})

	if err != nil {
		fmt.Println(err)
	}

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

	sCoords := make([]*services.Coordinate, len(city.Coords) / 20)


	for i, coord := range city.Coords {
		fmt.Println(i)
		if i % 20 == 0{
			sCoords[i / 20] = &services.Coordinate{
				Lat: coord.Lat,
				Lon: coord.Lon,
			}
		}

	}

	return rep.MongoCityToCity(&city)
}

func (rep *CityRepository) GetNearBy(longitude float64, latitude float64, limit int32) ([]*services.City, error) {
	fmt.Println("entry - get city by id")
	var cities []*City
	err := rep.mongo.Collection.Find(
		bson.M{"center":
					bson.M{"$near":
						bson.M{"$geometry":
							bson.M{"type":        "Point",
										"coordinates": []float64{longitude, latitude}}}}}).Limit(int(limit)).All(&cities)
	if err != nil {
		panic(err)
	}

	return rep.MongoCityListToCityList(cities), nil
}

func (rep *CityRepository) MongoCityListToCityList(cities []*City) ([]*services.City) {
	cCities := make([]*services.City, len(cities))
	for i, city := range cities {
		cCities[i] = rep.MongoCityToCity(city)
	}
	return cCities
}

func (rep *CityRepository) MongoCityToCity(city *City) (*services.City) {
	return &services.City{
		ID : city.ID,
		Coords: city.Coords,
		Center: &services.Coordinate{
			Lat: city.Center[0],
			Lon: city.Center[1],
		},
		Name : city.Name,
	}
}

func (rep *CityRepository) AnnounceActivity(cityid string, activityid string) error {
	_, err := rep.mongo.Collection.Upsert(bson.M{"cityid": cityid}, bson.M{"$addToSet": bson.M{"sportactivities":activityid}})

	return err
}

func (rep *CityRepository) GetAllCities() ([]*services.City, error) {
	var cities []*City
	rep.mongo.Collection.Find(bson.M{}).All(&cities)
	return rep.MongoCityListToCityList(cities),nil
}
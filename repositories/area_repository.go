package repositories

import (
	"fmt"
	"sportan/databases"
	"sportan/services"

	"gopkg.in/mgo.v2/bson"
)

type Area struct {
	Title       string                 `bson:"title,omitempty"`
	Description string                 `bson:"description,omitempty"`
	Coords      []*services.Coordinate `bson:"coords,omitempty"`
	Center      []float64              `bson:"center,omitempty"`
	ID          string                 `bson:"areaid,omitempty"`
}

type AreaRepository struct {
	mongo *databases.MongoConfig
}

func NewAreaRepository(mConfig *databases.MongoConfig) *AreaRepository {
	//initialize the area, if not exists ;-)
	//Collection is name

	return &AreaRepository{
		mongo: mConfig,
	}
}

func (rep *AreaRepository) CreateArea(id string, title string, description string, coordslist []*services.Coordinate, center *services.Coordinate) error {
	fmt.Println(title)
	//save center in database for faster access
	err := rep.mongo.Collection.Insert(
		&Area{
			ID:          id,
			Title:       title,
			Description: description,
			Center:      []float64{center.Lon, center.Lat},
			Coords:      coordslist})

	return err
}

func (rep *AreaRepository) GetNearBy(longitude float64, latitude float64, limit int32) (*[]Area, error) {
	var areas []Area
	err := rep.mongo.Collection.Find(
		bson.M{"center": bson.M{"$near": bson.M{"$geometry": bson.M{
			"type":        "Point",
			"coordinates": []float64{longitude, latitude}}}}}).Limit(int(limit)).All(&areas)
	if err != nil {
		panic(err)
	}

	return &areas, nil
}

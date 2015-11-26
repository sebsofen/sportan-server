package repositories

import (

	"sportan/databases"
	"sportan/services"

	"gopkg.in/mgo.v2/bson"

	"fmt"
)

type Area struct {
	Title       *string                 `bson:"title,omitempty"`
	Description *string                 `bson:"description,omitempty"`
	Coords      []*services.Coordinate `bson:"coords,omitempty"`
	Center      []float64              `bson:"center,omitempty"`
	ID          *string                 `bson:"areaid,omitempty"`
	CityId      *string             	   `bson:"cityid,omitempty"`
	Sports 		[]string				`bson:"sports,omitempty"`
	ImageId		*string					`bson:"iamgeid,omitempty"`
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



func (rep *AreaRepository) CreateArea(mArea *services.Area) {
	rep.mongo.Collection.Insert(
		rep.AreaToMongoArea(mArea))
}


func (rep *AreaRepository) AreaToMongoArea( mArea *services.Area) *Area {
	return &Area{
		ID:          mArea.ID,
		Title:       mArea.Title,
		Description: mArea.Description,
		Center:      []float64{mArea.Center.Lon, mArea.Center.Lat},
		Coords:      mArea.Coords,
		CityId:	     mArea.Cityid,
		ImageId: 	 mArea.Imageid,
		Sports:		 mArea.Sports,
	}
}



func (rep *AreaRepository) GetAllAreasInCity(cityid string) ([]*services.Area,error) {
	var areas []*Area
	rep.mongo.Collection.Find(bson.M{"cityid" : cityid}).All(&areas)

	return rep.MongoAreaListToAreaList(areas), nil
}


func (rep *AreaRepository) MongoAreaListToAreaList(areas []*Area) ([]*services.Area) {
	sAreas := make([]*services.Area, len(areas))
	for i, area := range areas {
		sAreas[i] = rep.MongoAreaToArea(area)
	}
	return sAreas
}

func (rep *AreaRepository) MongoAreaToArea(area *Area) (*services.Area) {
	return &services.Area{
		ID : area.ID,
		Coords: area.Coords,
		Center: &services.Coordinate{
			Lat: area.Center[1],
			Lon: area.Center[0],
		},
		Title : area.Title,
		Description: area.Description,
		Sports: area.Sports,
		Cityid:	     area.CityId,
		Imageid: 	area.ImageId,
	}
}

func (rep *AreaRepository) GetAreaById(areaid string) *services.Area {
	var area *Area
	rep.mongo.Collection.Find(bson.M{"areaid" : areaid}).One(&area)

	return rep.MongoAreaToArea(area)
}


func (rep *AreaRepository) UpdateArea(area *services.Area ) {
	fmt.Println("update area" + *area.ID)
	rep.mongo.Collection.Update(&Area{ID: area.ID}, rep.AreaToMongoArea(area))
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

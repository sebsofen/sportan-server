package repositories

import (

	"sportan/databases"
	"sportan/services"

	"gopkg.in/mgo.v2/bson"


	"gopkg.in/mgo.v2"
	"fmt"
)


type AreaRepository struct {
	mongo *databases.MongoConfig

}

func NewAreaRepository(mConfig *databases.MongoConfig) *AreaRepository {
	//initialize the area, if not exists ;-)
	//Collection is name
	/*
	index := mgo.Index{
		Key: []string{"$text:name", "$text:about"},
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	*/
//use sportan;
//	db.Areas.createIndex( { center: "2dsphere"});

	mConfig.Collection.EnsureIndex(
		mgo.Index{
			Key: []string{"$2dsphere:center"}})

	return &AreaRepository{
		mongo: mConfig,
	}
}



func (rep *AreaRepository) CreateArea(mArea *services.Area) error {
	fmt.Println(* mArea.Cityid)
	return rep.mongo.Collection.Insert(mArea)
}





func (rep *AreaRepository) GetAllAreasInCity(cityid string) ([]*services.Area,error) {
	var areas []*services.Area
	rep.mongo.Collection.Find(bson.M{"cityid" : cityid}).All(&areas)

	return areas, nil
}






func (rep *AreaRepository) GetAreaById(areaid string) *services.Area {
	var area *services.Area
	rep.mongo.Collection.Find(bson.M{"areaid" : areaid}).One(&area)

	return area
}


func (rep *AreaRepository) UpdateArea(area *services.Area ) error {
	return rep.mongo.Collection.Update(&services.Area{ID: area.ID}, area)
}









func (rep *AreaRepository) GetNearBy(longitude float64, latitude float64, limit int32) ([]*services.Area, error) {
	var areas []*services.Area
	err := rep.mongo.Collection.Find(
		bson.M{"center": bson.M{"$near": bson.M{"$geometry": bson.M{
			"type":        "Point",
			"coordinates": []float64{longitude, latitude}}}}}).Limit(int(limit)).All(&areas)
	if err != nil {
		panic(err)
	}

	return areas, nil
}

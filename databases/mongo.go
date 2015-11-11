package databases

import "gopkg.in/mgo.v2"

//MongoConfig store connection information
type MongoConfig struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

//NewMongoConfig Provide own configuration per collection
func NewMongoConfig(host string, database string, collection string) *MongoConfig {
	mSession, _ := mgo.Dial(host)
	c := mSession.DB(database).C(collection)
	return &MongoConfig{
		Session:    mSession,
		Collection: c,
	}
}

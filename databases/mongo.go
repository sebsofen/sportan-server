package databases

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

//MongoConfig store connection information
type MongoConfig struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

//NewMongoConfig Provide own configuration per collection
func NewMongoConfig(host []string, database string,  user string, pw string, collection string) (*MongoConfig, error) {
	fmt.Println("hoo" + user + " " + pw)
	mSession, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    host,
		Username: user,
		Password: pw,
		Database: database,
	})

	if err != nil {
		fmt.Println(err)
		return nil , err
	}



	c := mSession.DB(database).C(collection)
	fmt.Println("hii")
	return &MongoConfig{
		Session:    mSession,
		Collection: c,
	}, nil
}

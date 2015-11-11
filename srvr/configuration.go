package srvr

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration is used to store json information
type Configuration struct {
	OwnIP          string
	MongoHost      string
	MongoDatabase  string
	InfluxHost     string
	InfluxDatabase string
	InfluxPass     string
	InfluxUser     string
}

//NewConfiguration create configuration from file
func NewConfiguration(path string) *Configuration {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)
	c := Configuration{}
	err := decoder.Decode(&c)
	if err != nil {
		log.Fatal("malformed json file ", err)
	}
	return &c
}

package srvr

import (
	"fmt"
	"sportan/databases"
	"sportan/handlers"
	"sportan/repositories"
	"sportan/services"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type AppServer struct {
	server *thrift.TSimpleServer
	host   string
}

func NewAppServer(cfg Configuration) *AppServer {
	processor := thrift.NewTMultiplexedProcessor()

	//Register Processor after Processor
	//create config as u go
	fmt.Println("DATABASE " + cfg.MongoDatabase)
	metricsLogging := databases.NewMetricApi(cfg.InfluxHost, cfg.InfluxDatabase, cfg.InfluxUser, cfg.InfluxPass)

	//register User handler and stuff
	userCollection := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, "Users")
	userRepo := repositories.NewUserRepository(userCollection)
	userHandler := handlers.NewUserHandler(userRepo, metricsLogging)
	processor.RegisterProcessor("User", services.NewUserSvcProcessor(userHandler))

	//register Image handler and stuff
	imageCollection := databases.NewMongoConfig(cfg.MongoHost,cfg.MongoDatabase, "Images")
	imageRepo := repositories.NewImageRepository(imageCollection)
	imageHandler := handlers.NewImageHandler(imageRepo)
	processor.RegisterProcessor("Image", services.NewImageSvcProcessor(imageHandler))


	//register city handler and stuff
	cityCollection := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, "Cities")
	cityRepo := repositories.NewCityRepository(cityCollection)
	processor.RegisterProcessor("City", services.NewCitySvcProcessor(handlers.NewCityHandler(cityRepo, metricsLogging)))


	//register sporthandler and stuff
	sportCollection := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, "Sports")
	sportRepo := repositories.NewSportRepository(sportCollection,imageRepo)
	sportHandler := handlers.NewSportHandler(sportRepo,userRepo, imageRepo, metricsLogging)
	processor.RegisterProcessor("Sport", services.NewSportSvcProcessor(sportHandler))


	//register area handler and stuff
	areaCollection := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, "Areas")
	areaRepo := repositories.NewAreaRepository(areaCollection)
	processor.RegisterProcessor("Area", services.NewAreaSvcProcessor(handlers.NewAreaHandler(areaRepo, metricsLogging)))



	transport, err := thrift.NewTServerSocket(cfg.OwnIP)
	if err != nil {
		panic(err)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	return &AppServer{
		server: server,
		host:   cfg.OwnIP,
	}

}

func (ps *AppServer) Run() {
	fmt.Printf("server listening on %s\n", ps.host)
	err := ps.server.Serve()
	if err != nil {
		panic(err)
	}
}

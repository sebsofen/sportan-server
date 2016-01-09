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
	userCollection, err := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, cfg.MongoUser, cfg.MongoPw, "Users")
	userRepo := repositories.NewUserRepository(userCollection)
	userHandler := handlers.NewUserHandler(userRepo, metricsLogging)
	processor.RegisterProcessor(services.SERVICE_USER, services.NewUserSvcProcessor(userHandler))

	//register Image handler and stuff
	imageCollection, err := databases.NewMongoConfig(cfg.MongoHost,cfg.MongoDatabase, cfg.MongoUser, cfg.MongoPw, "Images")
	imageRepo := repositories.NewImageRepository(imageCollection)
	imageHandler := handlers.NewImageHandler(imageRepo,userRepo)
	processor.RegisterProcessor(services.SERVICE_IMAGE, services.NewImageSvcProcessor(imageHandler))


	//register city handler and stuff
	cityCollection, err := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, cfg.MongoUser, cfg.MongoPw, "Cities")
	cityRepo := repositories.NewCityRepository(cityCollection)
	processor.RegisterProcessor(services.SERVICE_CITY, services.NewCitySvcProcessor(handlers.NewCityHandler(cityRepo, metricsLogging)))


	//register sporthandler and stuff
	sportCollection, err := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, cfg.MongoUser, cfg.MongoPw, "Sports")
	sportRepo := repositories.NewSportRepository(sportCollection,imageRepo)
	sportHandler := handlers.NewSportHandler(sportRepo,userRepo, imageRepo, metricsLogging)
	processor.RegisterProcessor(services.SERVICE_SPORT, services.NewSportSvcProcessor(sportHandler))


	//register area handler and stuff
	areaCollection, err := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, cfg.MongoUser, cfg.MongoPw, "Areas")
	areaRepo := repositories.NewAreaRepository(areaCollection)
	processor.RegisterProcessor(services.SERVICE_AREA, services.NewAreaSvcProcessor(handlers.NewAreaHandler(areaRepo,  userRepo,metricsLogging)))


	//register sportActivity handler and stuff
	sportactivityCollection, err := databases.NewMongoConfig(cfg.MongoHost, cfg.MongoDatabase, cfg.MongoUser, cfg.MongoPw, "Activities")
	sportactivityRepo := repositories.NewSportActivityRepository(sportactivityCollection)
	sportactivityHandler := handlers.NewSportActivityHandler(sportactivityRepo,userRepo,areaRepo, cityRepo,metricsLogging)
	processor.RegisterProcessor(services.SERVICE_SPORTACTIVITY,services.NewSportActivitySvcProcessor(sportactivityHandler))

	if err != nil {
		panic(err)
	}

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

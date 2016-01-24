package repositories
import (
"sportan/databases"
"sportan/services"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/mgo.v2/bson"
)



type SportRepository struct {
	mongo *databases.MongoConfig
	repoI *ImageRepository
}


func NewSportRepository(mConfig *databases.MongoConfig, repoI *ImageRepository) *SportRepository {
	//Collection is name

	return &SportRepository{
		mongo: mConfig,
		repoI: repoI,
	}
}

func (rep *SportRepository) CreateSport(sport *services.Sport ) (*services.Sport){
	u, _ := uuid.NewV4()
	sportId := u.String()

	rep.mongo.Collection.Insert(sport)


	sport.ID = &sportId
	return sport

}

func (rep *SportRepository) GetAllSports() ([]*services.Sport) {
	var sports []*services.Sport
	rep.mongo.Collection.Find(bson.M{}).All(&sports)
	return sports
}

func (rep *SportRepository) GetSportById(id string) *services.Sport {
	var sport *services.Sport
	rep.mongo.Collection.Find(bson.M{"id": id}).One(&sport)
	return sport
}
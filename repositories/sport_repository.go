package repositories
import (
"sportan/databases"
"sportan/services"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Sport struct {
	Name *string `bson:"name,omitempty"`
	ID	*string `bson:"id,omitempty"`
	ImageId *string `bson:"imageid,omitempty"`
}

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
	rep.repoI.SaveImage(sport.Icon)
	//will update sport.Icon automatically.
	u, _ := uuid.NewV4()
	sportId := u.String()

	rep.mongo.Collection.Insert(
		&Sport {
			Name : sport.Name,
			ImageId : sport.Icon.ID,
			ID : &sportId,
		})


	sport.ID = &sportId
	return sport

}

func (rep *SportRepository) GetAllSports() ([]*services.Sport) {
	var sports []Sport
	rep.mongo.Collection.Find(bson.M{}).All(&sports)

	ssports := make([]*services.Sport, len(sports),len(sports))

	for i, sport := range sports {
		fmt.Println(&sport.ImageId)
		ssports[i] = &services.Sport{
			ID: sport.ID,
			Iconid: sport.ImageId,
			Name: sport.Name,
		}
	}

	return ssports
}
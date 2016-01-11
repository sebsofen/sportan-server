package repositories
import (
	"sportan/databases"
	"sportan/services"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/mgo.v2/bson"
)

type SportActivityRepository struct {
	mongo *databases.MongoConfig
}

func NewSportActivityRepository(mConfig *databases.MongoConfig) *SportActivityRepository {
	return &SportActivityRepository{
		mongo: mConfig,
	}
}


func (repo *SportActivityRepository) CreateSportActivity(activity *services.SportActivity) (*services.SportActivity, error) {

	if activity.GetID() == "" {
		u, _ := uuid.NewV4()
		id := u.String()
		activity.ID = &id
	}

	err := repo.mongo.Collection.Insert(activity)

	return activity, err

}

func (repo *SportActivityRepository) GetSportActivity(id string) (*services.SportActivity, error) {
	var activity *services.SportActivity
	err := repo.mongo.Collection.Find(bson.M{"activityid": id}).One(&activity)
	return activity, err
}

func (repo *SportActivityRepository) JoinUser(userid string, activityid string) error {
	_, err := repo.mongo.Collection.Upsert(bson.M{"activityid": activityid}, bson.M{"$addToSet": bson.M{"participants":userid}})
	return err
}



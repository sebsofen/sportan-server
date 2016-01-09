package repositories
import (
	"sportan/databases"
	"sportan/services"
	"github.com/nu7hatch/gouuid"
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

	if activity.GetID() == nil {
		u, _ := uuid.NewV4()
		activity.ID = u.String()
	}

	err := repo.mongo.Collection.Insert(activity)

	return activity, err

}

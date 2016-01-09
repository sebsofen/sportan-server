package handlers
import (
"sportan/repositories"
"sportan/databases"
"sportan/services"
)

type SportActivityHandler struct {
	userR   *repositories.UserRepository
	areaR   *repositories.AreaRepository
	cityR *repositories.CityRepository
	repo *repositories.SportActivityRepository
	metric *databases.MetricApi
}

func NewSportActivityHandler(repo *repositories.SportRepository,userR *repositories.UserRepository, areaR *repositories.AreaRepository, cityR *repositories.CityRepository, metricApi *databases.MetricApi) *SportActivityHandler {
	return &SportActivityHandler{
		repo:   repo,
		userR : userR,
		metric: metricApi,
	}



}

func (ch *SportActivityHandler) CreateActivity(token string, activity *services.SportActivity) (*services.SportActivity,error) {
	userid, _ := ch.userR.GetUserIdFromToken(token)
	host, err := ch.userR.GetFullUserInfo(userid)
	if err != nil {
		return err
	}

	activity.Hostid = userid

	//broadcast activity to all friends:
	mActivity, err := ch.repo.CreateSportActivity(activity)
	if err != nil {
		return err
	}


	//invite all friends to join activity
	friends := host.Friends
	if friends != nil {
		for _, friendid := range friends {
			//invite friend to activity
			ch.userR.AnnounceActivity(friendid,mActivity.ID)
		}
	}
	
	if activity.ActPublic {
		ch.cityR.AnnounceActivity(mActivity.Cityid,mActivity.ID)
	}



	return mActivity, err

}
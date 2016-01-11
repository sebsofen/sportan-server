package handlers
import (
"sportan/repositories"
"sportan/databases"
"sportan/services"
	"time"
)

type SportActivityHandler struct {
	userR   *repositories.UserRepository
	areaR   *repositories.AreaRepository
	cityR *repositories.CityRepository
	repo *repositories.SportActivityRepository
	metric *databases.MetricApi
}

func NewSportActivityHandler(repo *repositories.SportActivityRepository,userR *repositories.UserRepository, areaR *repositories.AreaRepository, cityR *repositories.CityRepository, metricApi *databases.MetricApi) *SportActivityHandler {
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
		return nil, err
	}

	activity.Hostid = &userid

	//broadcast activity to all friends:
	mActivity, err := ch.repo.CreateSportActivity(activity)
	if err != nil {
		return nil, err
	}


	//invite all friends to join activity
	friends := host.Friends
	if friends != nil {
		for _, friendid := range friends {
			//invite friend to activity
			ch.userR.AnnounceActivity(friendid,*(mActivity.ID))
		}
	}

	//Announce activity to user himself:
	ch.userR.AnnounceActivity(*(mActivity.Hostid),*(mActivity.ID))


	if activity.ActPublic != nil && *(activity.ActPublic) == true {
		ch.cityR.AnnounceActivity(*(mActivity.Cityid),*(mActivity.ID))
	}


	//finally add first participant to activity;
	//TODO this is disabled for debugging reasons
	//ch.repo.JoinUser(userid,*(mActivity.ID))
	return mActivity, err

}

func (ch *SportActivityHandler) GetAvailableActivityList(token string) ([]string, error) {
	activitylist := []string{}

	userid, err := ch.userR.GetUserIdFromToken(token)

	if err != nil {
		return nil, err
	}
	user, _ := ch.userR.GetFullUserInfo(userid)
	curTS := time.Now().UnixNano() / int64(time.Millisecond)
	if user.AnnouncedActivities != nil {
		for _, activityid := range user.AnnouncedActivities {
			sportActivity, err  := ch.repo.GetSportActivity(activityid)
			if err != nil || sportActivity.Date == nil || *(sportActivity.Date) < curTS{
				ch.userR.RemoveAnnouncedSportActivity(userid,activityid)
				continue
			}
			activitylist = append(activitylist,activityid)
		}
	}

	return activitylist, nil
}

func (ch *SportActivityHandler) GetActivity(token string, activityid string) (*services.SportActivity, error) {
	_, err := ch.userR.GetUserIdFromToken(token)
	if err != nil {
		return nil, err
	}

	activity, err := ch.repo.GetSportActivity(activityid)
	return activity, err
}

func (ch *SportActivityHandler) JoinActivity(token string, activityid string) error {

	curTS := time.Now().UnixNano() / int64(time.Millisecond)
	userid, err := ch.userR.GetUserIdFromToken(token)
	if err != nil{
		return err
	}
	user, err := ch.userR.GetFullUserInfo(userid)
	if err != nil {
		return err
	}

	if user.AnnouncedActivities != nil {
		isInvited := false
		for _, uActivityid := range user.AnnouncedActivities {
			if activityid == uActivityid {
				isInvited = true
			}
		}

		if isInvited {
			sportActivity, err := ch.repo.GetSportActivity(activityid)
			if err != nil {
				return err
			}

			if  sportActivity.Date != nil && *(sportActivity.Date) < curTS{
				if sportActivity.Participants == nil || len(sportActivity.Participants) <  int(*(sportActivity.MaxParticipants)) {
					ch.repo.JoinUser(userid,activityid)
				}
			}




		}
	}

	return nil
}


func (ch *SportActivityHandler) DeclineActivity(token string, activityid string) error {
	userid, err := ch.userR.GetUserIdFromToken(token)
	if err != nil{
		return err
	}
	err = ch.userR.RemoveAnnouncedSportActivity(userid, activityid)
	return err
}
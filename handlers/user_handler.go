package handlers

import (
	"fmt"
	"sportan/databases"
	"sportan/repositories"
	"sportan/services"

	"github.com/nu7hatch/gouuid"
	"time"
)

type UserHandler struct {
	repo   *repositories.UserRepository
	metric *databases.MetricApi
}

func NewUserHandler(mRepo *repositories.UserRepository, metricApi *databases.MetricApi) *UserHandler {
	return &UserHandler{
		repo:   mRepo,
		metric: metricApi,
	}
}

//Create and store user in database
//string is provided by device, username by server
func (ch *UserHandler) CreateUser(password string) (*services.UserCredentials, error) {
	//create userid
	u, _ := uuid.NewV4()
	userid := u.String()
	fmt.Println("[Info] New User")
	fmt.Println(userid)
	go ch.metric.PostDataPoint("new_users", 1)
	//add user to database
	ch.repo.AddUser(u.String(), password)
	return &services.UserCredentials{
		Identifier: userid,
	}, nil
}


func (ch *UserHandler) SetProfile(token string, profile *services.UserProfile)(error) {
	fmt.Println("Setting new Profile")
	_, err := ch.GetUserIdFromToken(token)

	if(err != nil){
		fmt.Println("cannot save profile, user not logged in!")
	}else{
		err := ch.repo.UpdateProfile(token,profile)
		if(err != nil){
			panic(err)
		}
		fmt.Println("profile updated")

	}
	return nil;

}


func (ch *UserHandler) RequestToken(username string, password string) (*services.ThriftToken,error){
	//create userid

	token := ch.repo.CreateTokenForUser(username,ch.repo.HashPw(password))
	return &services.ThriftToken{
		Token : token.Token,
		ValidityDuration : token.Validity - time.Now().UnixNano() / int64(time.Millisecond),
	}, nil
}

//useful function to retrieve userid for a given token, will return an error, if token is invalid
func (ch *UserHandler) GetUserIdFromToken(token string) (string,error){
	username, err := ch.repo.GetUserIdFromToken(token)
	fmt.Println(username)
	return username, err

}
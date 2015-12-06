package handlers

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"sportan/databases"
	"sportan/repositories"
	"sportan/services"
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

func (ch *UserHandler) GetMe(token string) (*services.User,error) {
	id, err := ch.GetUserIdFromToken(token)
	if err != nil {
		return nil, err
	}
	return ch.repo.GetFullUserInfo(id)
}

func (ch *UserHandler) SetProfile(token string, profile *services.Profile)(error) {
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
	return err;
}

func (ch *UserHandler) SetAdmin(token string, userid string) (error) {
	u, _ := ch.repo.GetUserByToken(token)
	fmt.Println(u.Role)
	//TODO CONTINUE HERE:
	//solve problem: idea is to extend services.user with issuperadmin function like in user_repository
	//wont be able to create function from outside package..
	return nil

}


func (ch *UserHandler) RequestToken(username string, password string) (*services.Token,error){
	//create userid

	token, err := ch.repo.CreateTokenForUser(username,ch.repo.HashPw(password))
	return token, err
}


//useful function to retrieve userid for a given token, will return an error, if token is invalid
func (ch *UserHandler) GetUserIdFromToken(token string) (string,error){
	username, err := ch.repo.GetUserIdFromToken(token)
	if &username == nil || username == "" || err != nil {
		err = services.NewInvalidToken()
	}
	return username, err

}
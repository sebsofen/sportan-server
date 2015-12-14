package handlers

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"sportan/databases"
	"sportan/repositories"
	"sportan/services"
	"sync"
	"errors"
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

//list<User> getFriends(1: string token);
func (ch *UserHandler) GetFriends(token string) ([]*services.User, error) {
	//TODO IMPLEMENT!!!
	id, err := ch.GetUserIdFromToken(token)
	if err != nil {
		return nil, err
	}

	user, _ := ch.repo.GetFullUserInfo(id)
	return ch.getUsersByStringList(user.Friends), nil

}

//void acceptFriendRequest(1: string token, 2: string userid);
func (ch *UserHandler) AcceptFriendRequest(token string, userid string) error {
	receiverid, err := ch.GetUserIdFromToken(token)
	if err != nil {
		return err
	}
	receiverUser, err := ch.repo.GetFullUserInfo(receiverid)
	for _, friendrequest := range receiverUser.Friendrequests {
		if friendrequest == userid {
			ch.repo.PutFriend(receiverid,userid)
			ch.repo.RemoveFriendRequest(receiverid,userid)
			return nil
		}
	}

	return errors.New("no friendrequest to accept")
}

func (ch *UserHandler) DeclineFriendRequest(token string, userid string) error {
	receiverid, err := ch.GetUserIdFromToken(token)
	if err != nil {
		return err
	}
	ch.repo.RemoveFriendRequest(receiverid,userid)
	return nil
}

func (ch *UserHandler) GetUserById(token string, userid string) (*services.User, error) {
	_, err := ch.GetUserIdFromToken(token)
	if err != nil {
		return nil, err
	}
	user, err := ch.repo.GetUserById(userid)
	//TODO DO REMOVE SOME UNNEEDED STUFF FROM USER!!!!!!!!!!!!!!!!!!!!!
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (ch *UserHandler) SendFriendRequest(token string, userid string) error {
	//save friend request in target user friendrequest
	receiverid, err := ch.GetUserIdFromToken(token)
	if err != nil {
		return err
	}

	//you can only send friendrequests to existing users
	_, err = ch.repo.GetUserById(userid)
	if err != nil {
		return err
	}

	//you cannot post new friendrequest to existing friend
	receiverUser, err := ch.repo.GetFullUserInfo(receiverid)
	for _, friend := range receiverUser.Friends {
		if friend == userid {
			return errors.New("users are friends already")
		}
	}

	return ch.repo.PutFriendRequest(receiverid, userid)

}

func (ch *UserHandler) getUsersByStringList(users []string) ([]*services.User) {
	var wg sync.WaitGroup
	sUsers := make([]*services.User, len(users))
	for i, userid := range users {
		wg.Add(1)
		go func(i int, userid string) {
			defer wg.Done()
			suser, _ := ch.repo.GetUserById(userid)
			sUsers[i] = suser
		}(i, userid)
	}

	wg.Wait()

	return sUsers
}

//list<User> getFriendRequests(1: string token);
func (ch *UserHandler) GetFriendRequests(token string) ([]*services.User, error) {
	//TODO IMPLEMENT!!!
	return nil,nil
}
//void requestFriend(1: string token, 2: string userid);
func (ch *UserHandler) RequestFriend(token string, userid string) error {
	//TODO IMPLEMENT!!!
	return nil
}
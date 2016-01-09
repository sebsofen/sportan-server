package repositories

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"sportan/databases"
	"time"
	"gopkg.in/mgo.v2/bson"
	"sportan/services"
	"strconv"
)

type UserRepository struct {
	mongo *databases.MongoConfig
}



//create a new user Repository
func NewUserRepository(mConfig *databases.MongoConfig) *UserRepository {

	return &UserRepository{
		mongo: mConfig,
	}
}

//Add user to database
func (rep *UserRepository) AddUser(uname string, userpassword string) error {
	hashpw := rep.HashPw(userpassword)
	err := rep.mongo.Collection.Insert(&services.User{Identifier: &uname, Password: &hashpw})
	return err
}

//generate token and set to current token in database
func (rep *UserRepository) CreateTokenForUser(uname string, hashedpw string) (*services.Token, error) {
	var ts int64
	ts =  24*3600*1000
	token := randToken()
	//update token in database for username
	tokenStruct := &services.Token{
		Token:    token,
		Validity: (time.Now().UnixNano()/1e6 + ts),
	}

	type M map[string]interface{}
	err := rep.mongo.Collection.Update(services.User{Identifier: &uname, Password: &hashedpw}, M{"$set": services.User{Token: tokenStruct}})
	return &services.Token{
		Token:    token,
		Validity: ts,
	}, err
}



func (rep *UserRepository) UpdateProfile(token string, profile *services.Profile) error {
	err := rep.mongo.Collection.Update(bson.M{"token.token": token},
	 bson.M{"$set" : services.User{Profile: profile }})

	return err
}

func (rep *UserRepository) GetUserIdFromToken(token string) (string, error) {
	user := services.User{}
	err := rep.mongo.Collection.Find(bson.M{"token.token": token}).One(&user)
	userid := ""
	if(err != nil){
		return "", err
	}
	userid = *user.Identifier
	return userid, err
}

//TODO : TO BE IMPLEMENTED
func (rep *UserRepository) GetUserById(userid string) (*services.User,error) {
	user := &services.User{}
	err := rep.mongo.Collection.Find(bson.M{"username": userid}).One(user)
	if err == nil {
		return user, nil
	}else {
		return nil,err
	}
}


func (rep *UserRepository) GetUserByToken(token string) (*services.User,error) {
	userid, err := rep.GetUserIdFromToken(token)
	if err != nil {
		return nil, err
	}
	return rep.GetUserById(userid)
}


func (rep *UserRepository) AnnounceActivity(receiverid string, activityid string) error {
	_, err := rep.mongo.Collection.Upsert(bson.M{"username": receiverid}, bson.M{"$addToSet": bson.M{"announced_activities":activityid}})
	return err
}


func(rep *UserRepository) PutFriendRequest(receiverid string, senderid string) error {
	_, err := rep.mongo.Collection.Upsert(bson.M{"username": receiverid}, bson.M{"$addToSet": bson.M{"friendrequests":senderid}})
	return err
}

func(rep *UserRepository) PutFriend(userid string, friendid string) error {
	_, err := rep.mongo.Collection.Upsert(bson.M{"username": userid}, bson.M{"$addToSet": bson.M{"friends":friendid}})
	return err
}

func (rep *UserRepository) RemoveFriend(userid string, friendid string) error {
	err := rep.mongo.Collection.Update(bson.M{"username": userid}, bson.M{"$pull": bson.M{"friends":friendid}})
	return err
}

func (rep *UserRepository) RemoveFriendRequest(userid string, friendid string) error {
	err := rep.mongo.Collection.Update(bson.M{"username": userid}, bson.M{"$pull": bson.M{"friendrequests":friendid}})
	return err
}

func (rep *UserRepository) IsAdmin(userid string) bool {
	var user *services.User
	rep.mongo.Collection.Find(bson.M{"username": userid}).One(&user)

	if user.Role != nil && *user.Role == services.ROLE_ADMIN || *user.Role == services.ROLE_SUPERADMIN {
		return true
	}
	return false
}

func (rep *UserRepository) GetFullUserInfo(userid string) (*services.User, error) {
	//TODO TO BE IMPLEMENTED
	var user *services.User

	err := rep.mongo.Collection.Find(bson.M{"username": userid}).One(&user)
	return user,err

}

func (rep *UserRepository) IsSuper(userid string) bool {
	var user *services.User
	rep.mongo.Collection.Find(bson.M{"username": userid}).One(&user)

	if user.Role != nil && *user.Role == services.ROLE_SUPERADMIN{
		return true
	}
	return false
}

func (rep *UserRepository) WasHere(userid string, areaid string, date int64) error {
	err := rep.mongo.Collection.Update(
		bson.M{"username" : userid},
		bson.M{"$addToSet": bson.M{"areasvisits": bson.M{strconv.FormatInt(date,10) : areaid}}})
	return err
}


//Set user to admin (or reset)
func (rep *UserRepository) SetAdmin(userid string, admin bool) error{
	if admin {
		admstr := services.ROLE_ADMIN
		err := rep.mongo.Collection.Update(services.User{Identifier: &userid}, bson.M{"$set": services.User{Role: &admstr}})
		return err
	}else {
		norole := ""
		err := rep.mongo.Collection.Update(services.User{Identifier: &userid}, bson.M{"$set": services.User{Role: &norole}})
		return err
	}
}


//hash password
func (rep *UserRepository) HashPw(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func randToken() string {
	size := 32
	rb := make([]byte, size)
	_, err := rand.Read(rb)

	if err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(rb)
}



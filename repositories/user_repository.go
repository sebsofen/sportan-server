package repositories

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sportan/databases"
	"time"

	"gopkg.in/mgo.v2/bson"
"sportan/services"
)

type UserRepository struct {
	mongo *databases.MongoConfig
}



const (
	ROLE_ADMIN ="admin"
	ROLE_SUPERADMIN = "superadmin"
)

// user as represented in database. be sure to update user conversion when needed
type User struct {
	Username     *string   `bson:"username,omitempty"`

	Password     *string   `bson:"password,omitempty"`
	Token        *Token   `bson:"token,omitempty"`
	Friends      []string `bson:"friends,omitempty"`
	Profile 	*Profile  `bson:"profile,omitempty"`
	Role 	*string `bson:"role,omitempty"`
}

type SUser struct {
	*services.User
}

func (us *SUser) ToMongoUser() (*User) {
	return &User {
		Username : us.Identifier,
		Role : us.Role,
	}
}

func (u *User) ToUser() (*services.User) {
	return &services.User{
		Identifier: u.Username,
		Role: u.Role,
	}
}

func (u *User) IsSuperAdmin() bool {
	return u.Role != nil && *u.Role == ROLE_SUPERADMIN
}



type Token struct {
	Token    string `bson:"token,omitempty"`
	Validity int64  `bson:"validity,omitempty"`
}

type Profile struct {
	Username    string `bson:"username,omitempty"`

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
	err := rep.mongo.Collection.Insert(&User{Username: &uname, Password: &hashpw})
	return err
}

//generate token and set to current token in database
func (rep *UserRepository) CreateTokenForUser(uname string, hashedpw string) *Token {
	ts := time.Now().UnixNano()/1e6 + 24*3600*1000
	token := randToken()
	//update token in database for username
	tokenStruct := &Token{
		Token:    token,
		Validity: ts,
	}
	fmt.Println(uname)
	type M map[string]interface{}

	err := rep.mongo.Collection.Update(User{Username: &uname, Password: &hashedpw}, M{"$set": User{Token: tokenStruct}})
	if err != nil {
		//this will only happen, if user does not exist or user/pw combi is wrong
		return &Token{
			Token:    "",
			Validity: 0,
		}
	}
	return tokenStruct
}



func (rep *UserRepository) UpdateProfile(token string, profile *services.UserProfile) error {
	err := rep.mongo.Collection.Update(bson.M{"token.token": token},
	 bson.M{"$set" : User{Profile: &Profile {
		Username : *profile.Username,
	}}})

	return err
}

func (rep *UserRepository) GetUserIdFromToken(token string) (string, error) {
	user := User{}
	err := rep.mongo.Collection.Find(bson.M{"token.token": token}).One(&user)
	userid := ""
	if(err != nil){
		fmt.Println("HEERE KOMMT EIN ERROR")
		panic(err)
	}
	userid = *user.Username
	return userid, err
}

//TODO : TO BE IMPLEMENTED
func (rep *UserRepository) GetUserById(userid string) (*services.User,error) {
	user := User{}
	err := rep.mongo.Collection.Find(bson.M{"username": userid}).One(&user)
	if err == nil {
		return user.ToUser(), nil
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

/*
//TODO Autocompletion for this
func (rep *UserRepository) MongoUserToUser(user *User) (*services.User) {
	return &services.User{
		Identifier: user.Username,
	}
}

//TODO Autocompletion for this
func (rep *UserRepository) UserToMongoUser(user *services.User) (*User) {
	return &User {
		Username: user.Identifier,
	}
}
*/

//TODO TO BE IMPLEMENTED
func (rep *UserRepository) IsAdmin(userid string) bool {
	var user *User
	rep.mongo.Collection.Find(bson.M{"username": userid}).One(&user)

	if user.Role != nil && *user.Role == ROLE_ADMIN || *user.Role == ROLE_SUPERADMIN {
		return true
	}
	return false
}



func (rep *UserRepository) IsSuper(userid string) bool {
	var user *User
	rep.mongo.Collection.Find(bson.M{"username": userid}).One(&user)

	if user.Role != nil && *user.Role == ROLE_SUPERADMIN{
		return true
	}
	return false
}

func (rep *UserRepository) SetAdmin(userid string) {

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

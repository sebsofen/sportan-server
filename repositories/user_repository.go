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

type User struct {
	Username     string   `bson:"username,omitempty"`
	Password     string   `bson:"password,omitempty"`
	Token        *Token   `bson:"token,omitempty"`
	Friends      []string `bson:"friends,omitempty"`
	Profile 	*Profile  `bson:"profile,omitempty"`
}

type Token struct {
	Token    string `bson:"token,omitempty"`
	Validity int64  `bson:"validity,omitempty"`
}

type Profile struct {
	Username    string `bson:"username,omitempty"`

}

func NewUserRepository(mConfig *databases.MongoConfig) *UserRepository {
	return &UserRepository{
		mongo: mConfig,
	}
}

//Add user to database
func (rep *UserRepository) AddUser(uname string, userpassword string) error {
	err := rep.mongo.Collection.Insert(&User{Username: uname, Password: rep.HashPw(userpassword)})
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

	err := rep.mongo.Collection.Update(User{Username: uname, Password: hashedpw}, M{"$set": User{Token: tokenStruct}})
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
	userid = user.Username
	return userid, err
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

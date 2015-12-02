package models
import (
	"sportan/services"

)

type SUser struct {
	*services.User
}
// user as represented in database. be sure to update user conversion when needed
type User struct {
	Username     *string   `bson:"username,omitempty"`
	Password     *string   `bson:"password,omitempty"`
	Token        *Token   `bson:"token,omitempty"`
	Friends      []string `bson:"friends,omitempty"`
	Profile 	*Profile  `bson:"profile,omitempty"`
	Role 	*string `bson:"role,omitempty"`
}

type Token struct {
	Token    string `bson:"token,omitempty"`
	Validity int64  `bson:"validity,omitempty"`
}

type Profile struct {
	Username    string `bson:"username,omitempty"`

}

func (u *User) ToUser() (*services.User) {
	return &services.User{
		Identifier: u.Username,
		Role: u.Role,
	}
}

func (su *SUser) ToMongoUser() (*User) {
	return &User{
		Username: su.Identifier,
		Role : su.Role,
	}
}
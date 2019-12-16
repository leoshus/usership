package model

import (
	"fmt"
	"usership/db"
)

type User struct {
	Id       int64 `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"type"`
}
//format User when print
func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s>",u.Id,u.Name,u.UserType)
}

//ListUsers: return all users list in database
func ListUsers() ([]User, error) {
	users := make([]User, 0)
	err := db.DB().Model(&users).Select()
	return users, err
}

//InsertUser: store user data to table user
func InsertUser(u *User) error {
	u.UserType = "user"
	err := db.DB().Insert(u)
	return err
}
//ExistUser: check user is exist by userId
func ExistUser(userId int64) error {
	u := new(User)
	err := db.DB().Model(u).Where("id=?", userId).Select()
	return err
}
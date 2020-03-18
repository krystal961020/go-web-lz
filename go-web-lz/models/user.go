package models

import "go-web-lz/datasource"

type User struct {
	datasource.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func GetUser(username, password string) bool {
	var user User
	datasource.GetDB().Select("id").Where(User{
		UserName: username,
		Password: password,
	}).First(&user)
	return user.ID > 0

}

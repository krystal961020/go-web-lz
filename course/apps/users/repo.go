package users

import "course/datasource"

type UserRepository interface {
	CheckUser(username string, password string) (user UserInfo)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

//登录
func (n userRepository) CheckUser(username string, password string) (user UserInfo) {
	db := datasource.GetDB()
	db.Where("username = ? and password = ?", username, password).First(&user)
	return
}

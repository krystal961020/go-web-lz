package users

import (
	"course/middleware"
	"course/utils"
	"fmt"
)

type UserService interface {
	Login(m map[string]string) (result utils.Result)
}
type userServices struct {
}

func NewUserServices() UserService {
	return &userServices{}
}

var userRepo = NewUserRepository()

func (u userServices) Login(m map[string]string) (result utils.Result) {

	if m["username"] == "" {
		result.Code = 1000
		result.Msg = "请输入用户名！"
		return
	}
	if m["password"] == "" {
		result.Code = 1000
		result.Msg = "请输入密码！"
		return
	}
	user := userRepo.CheckUser(m["username"], m["password"])
	if user.ID == 0 {
		result.Code = 1000
		result.Msg = "用户名或密码错误!"
		return
	}
	token := middleware.GenerateToken(user.Username, user.ID)
	data := map[string]string{
		"token": token,
	}
	a := middleware.GetUserID(token)
	fmt.Println(a)
	result.Code = 1001
	result.Data = data
	result.Msg = "登录成功"
	return
}

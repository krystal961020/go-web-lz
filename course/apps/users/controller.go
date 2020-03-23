package users

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx     iris.Context
	Service UserService
}

func NewUserController() *UserController {
	return &UserController{Service: NewUserServices()}
}

func (u *UserController) Login(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	fmt.Println(password)
	fmt.Println("password")
	m := make(map[string]string)
	m["username"] = username
	m["password"] = password
	result := u.Service.Login(m)
	ctx.JSON(result)

}

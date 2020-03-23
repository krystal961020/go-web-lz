package route

import (
	"course/apps/users"
	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {

	// Simple group: v1.
	v1 := app.Party("/api/v1")

	{
		userHandler := users.NewUserController()
		v1.Post("/login", userHandler.Login)
	}

}

package routers

import (
	beego "github.com/beego/beego/v2/server/web"

	"beego-poc/controllers"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(routes),
)

func routes(userController *controllers.UserController) {
	beego.Router("/user/:id", userController, "get:Get")
}

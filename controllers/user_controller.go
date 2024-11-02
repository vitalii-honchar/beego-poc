package controllers

import (
	"beego-poc/services"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) Get() {
	logs.Info("Get user by ID: svc = %+v", u.UserService)
	idStr := u.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Ctx.WriteString("Invalid ID")
		return
	}
	user := u.UserService.GetByID(id)

	u.Data["User"] = user
	u.TplName = "pages/user/profile.html"
}

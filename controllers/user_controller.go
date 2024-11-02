package controllers

import (
	"beego-poc/models"
	"beego-poc/services"
	"strconv"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/client/orm"
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

	o := orm.NewOrm()
	user := &models.User{ID: id}
	if err := o.Read(user); err != nil {
		u.Ctx.WriteString("User not found")
		return
	}

	u.Data["User"] = user
	u.TplName = "pages/user/profile.html"
}

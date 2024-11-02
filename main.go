package main

import (
	"beego-poc/controllers"
	"beego-poc/routers"
	"beego-poc/services"
	"context"

	beego "github.com/beego/beego/v2/server/web"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		routers.Module,
		controllers.Module,
		services.Module,
		fx.Invoke(RunBeego),
	).Run()
}

func RunBeego(lc fx.Lifecycle) {
	beego.BConfig.Listen.EnableAdmin = true
	beego.BConfig.Listen.AdminAddr = "localhost"
	beego.BConfig.Listen.AdminPort = 8088
	// beego.BConfig.WebConfig.AutoRender = true
	// beego.SetViewsPath("views")

	beego.BConfig.RunMode = "dev"

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			go beego.Run()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			beego.BeeApp.Server.Shutdown(ctx)
			return nil
		},
	})
}

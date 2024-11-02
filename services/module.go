package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(fx.Annotate(newUserService, fx.As(new(UserService)))),
)

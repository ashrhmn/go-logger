package user

import (
	"github.com/ashrhmn/go-logger/utils"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"users",
	fx.Provide(
		newUserService,
		utils.ProvideController(newUserController),
	),
)

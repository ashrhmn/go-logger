package auth

import (
	"github.com/ashrhmn/go-logger/utils"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"auth",
	fx.Provide(
		newAuthService,
		utils.ProvideController(newAuthController),
	),
)

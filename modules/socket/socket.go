package socket

import (
	"github.com/ashrhmn/go-logger/utils"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"socket",
	fx.Provide(
		utils.ProvideController(newSocketController),
	),
)

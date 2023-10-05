package utils

import (
	"github.com/ashrhmn/go-logger/types"
	"go.uber.org/fx"
)

func ProvideController(c any) interface{} {
	return fx.Annotate(
		c,
		fx.As(new(types.Controller)),
		fx.ResultTags(`group:"controllers"`),
	)
}

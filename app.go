package main

import (
	"github.com/ashrhmn/go-logger/modules"
	"github.com/ashrhmn/go-logger/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(server.Module, modules.Register, fx.NopLogger).Run()
}

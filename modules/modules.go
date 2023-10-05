package modules

import (
	"github.com/ashrhmn/go-logger/modules/auth"
	"github.com/ashrhmn/go-logger/modules/logging"
	"github.com/ashrhmn/go-logger/modules/storage"
	"github.com/ashrhmn/go-logger/modules/user"
	"go.uber.org/fx"
)

var Register = fx.Module("modules",
	storage.Module,
	logging.Module,
	user.Module,
	auth.Module,
)

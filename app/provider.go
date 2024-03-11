package app

import (
	"app/app/controller"
	"app/app/service"

	"github.com/google/wire"
)

var Provider = wire.NewSet(controller.Provider, service.Provider)

package controller

import (
	"app/infra"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewBaseCtl, wire.Bind(new(infra.Controller), new(*BaseCtl)),
)

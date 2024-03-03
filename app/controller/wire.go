package controller

import (
	"app/infra"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	wire.Struct(new(BaseCtl), "*"), wire.Bind(new(infra.Controller), new(*BaseCtl)),
)

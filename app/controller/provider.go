package controller

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	wire.Struct(new(BaseCtl), "*"),
	wire.Struct(new(UserCtl), "*"),
)

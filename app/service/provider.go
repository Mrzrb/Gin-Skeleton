package service

import (
	"app/app/service/impl"

	"github.com/google/wire"
)

var Provider = wire.NewSet(
	wire.Struct(new(impl.UserService), "*"),
	wire.Bind(new(UserService), new(*impl.UserService)),
)

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package wires

import (
	"app/app"
	"app/infra"

	"github.com/google/wire"
)

func InitApp() (*app.App, error) {
	panic(wire.Build(wire.Struct(new(app.App), "*"), app.Provider, infra.Provider))
}

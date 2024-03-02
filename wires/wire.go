//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package wires

import (
	"app/infra"

	"github.com/google/wire"
)

func InitApp() (*infra.App, error) {
	panic(wire.Build(infra.NewApp))
}

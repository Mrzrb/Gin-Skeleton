package infra

import "github.com/google/wire"

var Provider = wire.NewSet(NewCommonDB)

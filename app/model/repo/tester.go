package repo

import "gorm.io/gen"

type ITest interface {
	// select * from @@table where name = @name
	GetByName(name string) ([]gen.T, error)
}

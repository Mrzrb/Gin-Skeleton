package repo

import (
	"gorm.io/gen"
	"gorm.io/gorm/schema"
)

type IBase[T any] interface {
	gen.SubQuery
	Create(values ...*T) error
	CreateInBatches(values []*T, batchSize int) error
	Save(values ...*T) error
	First() (*T, error)
	Take() (*T, error)
	Last() (*T, error)
	Find() ([]*T, error)
	schema.Tabler
}

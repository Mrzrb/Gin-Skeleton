package dao

import (
	"app/infra"
	"database/sql/driver"
)

type Password string

func (p Password) Value() (driver.Value, error) {
	encrypted, err := infra.PasswordHash(string(p))
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

type Model struct {
	Id        uint  `gorm:"id;primarykey"`
	CreatedAt int64 `gorm:"column:createdAt"`
	UpdatedAt int64 `gorm:"column:updatedAt"`
}

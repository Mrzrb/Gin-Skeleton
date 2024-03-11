package dao

import "gorm.io/gorm"

type Demo struct {
	Name string
	gorm.Model
}

package infra

import (
	"app/helpers"

	"gorm.io/gen/examples/dal/query"
	"gorm.io/gorm"
)

type CommonDB *gorm.DB

func NewCommonDB() CommonDB {
	cmDb := CommonDB(helpers.MysqlClients[helpers.CommonMysqlClient])
	query.SetDefault(cmDb)
	return cmDb
}

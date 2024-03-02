package infra

import (
	"app/app/model/dao"
	"app/helpers"

	"gorm.io/gorm"
)

type CommonDB *gorm.DB

func NewCommonDB() CommonDB {
	cmDb := CommonDB(helpers.MysqlClients[helpers.CommonMysqlClient])
	dao.SetDefault(cmDb)
	return cmDb
}

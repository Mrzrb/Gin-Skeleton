package helpers

import (
	"app/config"
	"app/infra"

	"gorm.io/gorm"
)

var MysqlClients map[string]*gorm.DB

func InitMysql() {
	for k, mc := range config.Conf.Mysql {
		client, err := infra.InitMysqlClient(mc)
		if err != nil {
			panic(err)
		}
		MysqlClients[k] = client
	}
}

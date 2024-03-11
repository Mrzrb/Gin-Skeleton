package helpers

import (
	"app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlClients = map[string]*gorm.DB{}

const (
	CommonMysqlClient = "base"
)

func InitMysql() {
	for k, mc := range config.Conf.Mysql {
		client, err := initMysqlClientt(mc)
		if err != nil {
			panic(err)
		}
		MysqlClients[k] = client
	}
}

func initMysqlClientt(conf config.MysqlConf) (client *gorm.DB, err error) {
	conf.CheckConf()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s&parseTime=True&loc=Local&interpolateParams=%s",
		conf.User,
		conf.Password,
		conf.Addr,
		conf.DataBase,
		conf.ConnTimeOut,
		conf.ReadTimeOut,
		conf.WriteTimeOut,
		conf.InterpolateParams,
	)

	if conf.Charset != "" {
		dsn = dsn + "&charset=" + conf.Charset
	}

	c := &gorm.Config{
		SkipDefaultTransaction:                   true,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		AllowGlobalUpdate:                        false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}

	client, err = gorm.Open(mysql.Open(dsn), c)
	if err != nil {
		return client, err
	}

	sqlDB, err := client.DB()
	if err != nil {
		return client, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(conf.ConnMaxLifeTime)

	// only for go version >= 1.15 设置最大空闲连接时间
	sqlDB.SetConnMaxIdleTime(conf.ConnMaxIdlTime)

	return client, nil
}

package tests

import (
	"app/app"
	"app/app/model/query"
	"app/config"
	"app/helpers"
	"app/wires"

	"github.com/gin-gonic/gin"
)

var (
	TestApp *app.App
	Q       *query.Query
)

func init() {
	config.InitConf("unit")
	engine := gin.New()
	helpers.Init(engine)
	app, err := wires.InitApp()
	if err != nil {
		panic(err)
	}
	query.SetDefault(helpers.MysqlClients[helpers.CommonMysqlClient])
	TestApp = app
}

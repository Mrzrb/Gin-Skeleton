package tests

import (
	"app/config"
	"app/helpers"
	"app/infra"
	"app/wires"

	"github.com/gin-gonic/gin"
)

var TestApp *infra.App

func init() {
	config.InitConf("unit")
	engine := gin.New()
	helpers.Init(engine)
	app, err := wires.InitApp()
	if err != nil {
		panic(err)
	}
	TestApp = app
}

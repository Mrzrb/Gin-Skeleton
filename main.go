package main

import (
	"app/config"
	"app/helpers"
	"app/infra"
	"app/tools"
	"app/wires"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConf("dev")
	engine := gin.New()
	helpers.Init(engine)
	app, err := wires.InitApp()
	if err != nil {
		panic(err)
	}
	setUpRouter(engine, app)
	tools.GenerateDoc(engine)
	helpers.HttpServer(engine)
}

func setUpRouter(engine *gin.Engine, app *infra.App) error {
	engine.Use(cors.Default())
	for _, c := range app.Ctls {
		c.Routes(engine)
	}
	return nil
}

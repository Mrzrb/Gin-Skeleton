package main

import (
	"app/config"
	"app/helpers"
	"app/wires"

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
	app.Run(engine)
}

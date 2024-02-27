package main

import (
	"app/config"
	"app/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConf("dev")
	engine := gin.New()
	helpers.Init(engine)
	helpers.HttpServer(engine)
}

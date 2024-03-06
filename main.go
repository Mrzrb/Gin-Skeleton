package main

import (
	"app/config"
	"app/helpers"
	"app/infra"
	"app/wires"

	"github.com/Mrzrb/astra"
	"github.com/Mrzrb/astra/inputs"
	"github.com/Mrzrb/astra/outputs"
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
	generateDoc(engine)
	helpers.HttpServer(engine)
}

func setUpRouter(engine *gin.Engine, app *infra.App) error {
	for _, c := range app.Ctls {
		c.Routes(engine)
	}
	return nil
}

func generateDoc(engine *gin.Engine) error {
	gen := astra.New(inputs.WithGinInput(engine), outputs.WithOpenAPIOutput("openapi.generated.yaml"))
	config := astra.Config{
		Title:   "Example API",
		Version: "1.0.0",
		Host:    "localhost",
		Port:    9000,
	}
	gen.SetConfig(&config)
	err := gen.Parse()
	if err != nil {
		panic(err)
	}

	return err
}

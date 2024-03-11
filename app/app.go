package app

import (
	"app/app/controller"
	"app/helpers"
	"app/infra"
	"app/tools"
	"reflect"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	BaseCtl controller.BaseCtl
	UserCtl controller.UserCtl
	Ctls    []infra.Controller `wire:"-"`
}

func NewApp(baseCtl *controller.BaseCtl) *App {
	app := &App{
		Ctls: []infra.Controller{baseCtl},
	}

	return app
}

func (a *App) Run(engine *gin.Engine) {
	v := reflect.ValueOf(a).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Addr()
		if fieldValue := field.Interface(); fieldValue != nil {
			if c, ok := fieldValue.(infra.Controller); ok {
				a.Ctls = append(a.Ctls, c)
			}
		}
	}
	a.setUpRouter(engine)
	go tools.GenerateDoc(engine)
	helpers.HttpServer(engine)
}

func (a *App) setUpRouter(engine *gin.Engine) error {
	engine.Use(cors.Default())
	for _, c := range a.Ctls {
		c.Routes(engine)
	}
	return nil
}

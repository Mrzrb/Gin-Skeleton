package controller

import (
	"app/infra"

	"github.com/gin-gonic/gin"
)

type BaseCtl struct{}

// Routes implements infra.Controller.
func (b *BaseCtl) Routes(engine *gin.Engine) {
	engine.POST("/base", func(ctx *gin.Context) {})
}

func NewBaseCtl() *BaseCtl {
	return &BaseCtl{}
}

var _ infra.Controller = (*BaseCtl)(nil)

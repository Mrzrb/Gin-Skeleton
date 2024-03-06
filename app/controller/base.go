package controller

import (
	"app/infra"
	"app/infra/messages"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseCtl struct {
	EmailSrv messages.EmailProvider
}

type Demo struct {
	AAd string `json:"aAd"`
}

func FnWarp(fn infra.ControllerMethod) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rsp, err := fn(ctx)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusOK, rsp)
	}
}

// Routes implements infra.Controller.
func (b *BaseCtl) Routes(engine *gin.Engine) {
}

var _ infra.Controller = (*BaseCtl)(nil)

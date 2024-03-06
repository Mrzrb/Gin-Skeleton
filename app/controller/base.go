package controller

import (
	"app/infra"
	"app/infra/messages"
	"app/tools"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseCtl struct {
	EmailSrv messages.EmailProvider
}

type Demo struct {
	AAd string `json:"aAd" bindings:required`
}

type Resp struct {
	Avatar   string
	Grade    int
	ID       int
	NickName string
	Phone    string
	PlayRole int
	SchoID   string
	School   string
	Sex      int
	UID      int
	UName    []string
	Hit      int
	WriteHit int
}

func (d *BaseCtl) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Resp{})
}

func (d *BaseCtl) Hellp(ctx *gin.Context) {
	var req Demo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.RenderJsonFail(ctx, errors.New("test"))
		return
	}
	tools.RenderJsonSucc(ctx, Resp{})
}

// Routes implements infra.Controller.
func (b *BaseCtl) Routes(engine *gin.Engine) {
	engine.POST("/test", b.Hellp)
}

var _ infra.Controller = (*BaseCtl)(nil)

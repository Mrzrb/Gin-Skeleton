package controller

import (
	"app/infra"
	"app/infra/messages"
	"app/tools"
	"errors"

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
	Rec      *Resp
}

// 生成unsafe rust代码
// go run github.com/owenthereal/unsafe go2rust -i app/controller/base.go -o app/controller/base.rs
func (d *BaseCtl) GetUserInfo(ctx *gin.Context) {
	var req Demo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.RenderJsonFail(ctx, errors.New("test"))
		return
	}
	tools.RenderJsonSucc(ctx, Resp{})
}

// Routes implements infra.Controller.
func (b *BaseCtl) Routes(engine *gin.Engine) {
	engine.POST("/getUserInfo", b.GetUserInfo)
}

var _ infra.Controller = (*BaseCtl)(nil)

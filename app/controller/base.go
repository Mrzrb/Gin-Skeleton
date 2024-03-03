package controller

import (
	"app/config"
	"app/infra"
	"app/infra/messages"

	"github.com/gin-gonic/gin"
)

type BaseCtl struct {
	EmailSrv messages.EmailProvider
}

// Routes implements infra.Controller.
func (b *BaseCtl) Routes(engine *gin.Engine) {
	engine.POST("/base", func(ctx *gin.Context) {
		client := b.EmailSrv
		m := messages.NewEmail(config.Conf.EmailNoreply.EmailUser, "test send email", "<h1>This</h1> is a test email from go skeleton", []string{"mrzhangrb@outlook.com"})
		err := client.Send(m)
		ctx.JSON(0, err)
	})
}

var _ infra.Controller = (*BaseCtl)(nil)

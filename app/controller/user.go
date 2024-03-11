package controller

import (
	"app/app/dto/user"
	"app/app/service"
	"app/app/types"
	"app/tools"

	"github.com/gin-gonic/gin"
)

type UserCtl struct {
	UserSrv service.UserService
}

func (u *UserCtl) Routes(engine *gin.Engine) {
	group := engine.Group("/user")
	group.POST("/email/send-verification-code", u.EmailSendVerificationCode)
}

func (u *UserCtl) EmailSendVerificationCode(ctx *gin.Context) {
	var req user.EmailSendVerificationCodeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.RenderJsonFail(ctx, err)
		return
	}
	err := u.UserSrv.EmailSendVerificationCode(ctx, types.Email(req.Email))
	if err != nil {
		tools.RenderJsonFail(ctx, err)
		return
	}
	tools.RenderJsonSucc(ctx, nil)
}

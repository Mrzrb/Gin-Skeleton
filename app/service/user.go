package service

import (
	"app/app/types"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	EmailSendVerificationCode(ctx *gin.Context, email types.Email) error
	EmailVerification(ctx *gin.Context, email types.Email, code string) error
}

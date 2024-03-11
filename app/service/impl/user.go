package impl

import (
	"app/app/consts/errors"
	"app/app/types"
	"app/infra"
	"app/infra/messages"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	Email messages.EmailProvider
	Cache *infra.Cache
	Idg   *infra.IdGenerator
}

func (u *UserService) EmailSendVerificationCode(ctx *gin.Context, email types.Email) error {
	code := u.Idg.GenerateVerificationCode(6)
	msg := messages.NewEmail("VerificationCode", code, []string{string(email)})
	err := u.Email.Send(msg)
	if err != nil {
		return errors.ServerError.Wrap(err)
	}
	_, err = u.Cache.Proxy().Set(ctx, "verification:"+email.String(), code, time.Minute*5).Result()
	if err != nil {
		return errors.ServerError.Wrap(err)
	}

	return nil
}

func (u *UserService) EmailVerification(ctx *gin.Context, email types.Email, code string) error {
	panic("not implemented") // TODO: Implement
}

package user

type EmailSendVerificationCodeReq struct {
	Email string `binding:"required,email" json:"email"`
}

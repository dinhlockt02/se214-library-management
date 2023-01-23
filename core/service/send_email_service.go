package coreservice

type SendEmailService interface {
	SendResetPasswordMail(name string, email string, resetCode int) error
}

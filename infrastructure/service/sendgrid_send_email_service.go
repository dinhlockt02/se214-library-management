package service

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	coreservice "daijoubuteam.xyz/se214-library-management/core/service"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

type SendGridSendEmailService struct {
}

func NewSendGridSendEmailService() coreservice.SendEmailService {
	return SendGridSendEmailService{}
}

func (s SendGridSendEmailService) SendResetPasswordMail(name string, email string, resetCode int) error {
	from := mail.NewEmail("Library Management", "dinhlockt02@gmail.com")
	subject := "Forget password requested"
	to := mail.NewEmail(name, email)
	htmlContent := fmt.Sprintf("Your reset password code is <strong>%d<strong>", resetCode)
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		return coreerror.NewInternalServerError("Send email failed", err)
	}
	return nil
}

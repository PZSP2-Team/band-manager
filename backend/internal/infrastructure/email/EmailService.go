package email

import (
	"band-manager-backend/internal/model"
	"fmt"
	"net/smtp"
	"os"
)

type EmailService struct {
	from     string
	password string
	smtpHost string
	smtpPort string
}

func NewEmailService() *EmailService {
	return &EmailService{
		from:     os.Getenv("EMAIL_FROM"),
		password: os.Getenv("APP_PASSWORD"),
		smtpHost: os.Getenv("SMTP_HOST"),
		smtpPort: os.Getenv("SMTP_PORT"),
	}
}

func (s *EmailService) SendEventEmail(event *model.Event, recipients []*model.User) error {
	fmt.Printf("Config: host=%s, port=%s, from=%s\n", s.smtpHost, s.smtpPort, s.from)

	if len(recipients) == 0 {
		return nil
	}

	auth := smtp.PlainAuth("", s.from, s.password, s.smtpHost)
	fmt.Printf("Sending to: %v\n", recipients[0].Email)
	subject := fmt.Sprintf("Nowe wydarzenie: %s", event.Title)
	body := fmt.Sprintf(
		"Nazwa: %s\nOpis: %s\nMiejsce: %s\nData: %s",
		event.Title,
		event.Description,
		event.Location,
		event.Date.Format("02.01.2006 15:04"),
	)

	for _, recipient := range recipients {
		msg := []byte(fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\n\r\n%s",
			recipient.Email,
			s.from,
			subject,
			body,
		))

		if err := smtp.SendMail(
			s.smtpHost+":"+s.smtpPort,
			auth,
			s.from,
			[]string{recipient.Email},
			msg,
		); err != nil {

			fmt.Printf("Błąd wysyłania maila do %s: %v\n", recipient.Email, err)
		}
	}

	return nil
}

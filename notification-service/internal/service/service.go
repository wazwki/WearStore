package service

import (
	"strconv"
	"time"

	"github.com/wazwki/WearStore/notification-service/pkg/metrics"
	"gopkg.in/gomail.v2"
)

type ServiceInterface interface {
	SendNotification(emails []string) error
}

type Service struct {
	SMTPConfig map[string]string
}

func NewService(config map[string]string) ServiceInterface {
	return &Service{SMTPConfig: config}
}

func (s *Service) SendNotification(emails []string) error {
	start := time.Now()
	message := gomail.NewMessage()

	message.SetHeader("From", s.SMTPConfig["from"])
	message.SetHeader("To", emails...)
	message.SetHeader("Subject", "New order")

	message.SetBody("text/html", `
        <html>
            <body>
                <h1>Thank you for your order!</h1>
                <p><b>Hello!</b> Check your order in profile ^_^</p>
                <p>Thanks,<br>WearStore</p>
            </body>
        </html>
    `)

	port, err := strconv.Atoi(s.SMTPConfig["port"])
	if err != nil {
		return err
	}
	dialer := gomail.NewDialer(s.SMTPConfig["host"], port, s.SMTPConfig["username"], s.SMTPConfig["password"])

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	metrics.ServiceDuration.WithLabelValues("notification-service.SendNotification").Observe(time.Since(start).Seconds())
	return nil
}

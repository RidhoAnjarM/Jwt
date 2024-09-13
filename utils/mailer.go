package utils

import (
	"os"

	"github.com/nothing2512/mailer"
)

func SendMail(title, target, view string, data interface{}) error {
	m, err := mailer.Init(
		os.Getenv("MAIL_USER"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_HOST"),
		os.Getenv("MAIL_PORT"),
	)
	if err != nil {
		return err
	}

	m.Recipients(target)
	m.Subject(title)
	m.SetHTMLFile(view, data)

	return m.Send()
}

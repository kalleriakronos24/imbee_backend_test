package mail

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type tCarbonCopy struct {
	address string
	name    string
}

type tSendMail struct {
	from        []string
	mailTo      []string
	cc          []*tCarbonCopy
	subject     string
	body        string
	template    string
	attachments string
}

func SendEmail(t *tSendMail) *gomail.Message {

	mail := gomail.NewMessage()
	mail.SetHeader("From", t.from...)
	mail.SetHeader("To", t.mailTo...)

	if len(t.cc) > 0 {
		for _, v := range t.cc {
			mail.SetAddressHeader("Cc", v.address, v.name)
		}
	}
	mail.SetHeader("Subject", t.subject)
	mail.SetBody(t.body, t.template)
	mail.Attach(t.attachments)

	d := gomail.NewDialer(viper.GetString("MAIL_HOST"), viper.GetInt("MAIL_PORT"), viper.GetString("MAIL_USERNAME"), viper.GetString("MAIL_PASSWORD"))

	if err := d.DialAndSend(mail); err != nil {
		panic(err)
	}

	return mail
}

package utils

import "net/mail"

func EmailFormatValidation(email string) (err error) {
	_, err = mail.ParseAddress(email)
	return err
}

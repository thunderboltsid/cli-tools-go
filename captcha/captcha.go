package captcha

import (
	"io"
)

type Captcha interface {
	ConfirmPhrase(reader io.Reader) error
}

func New() Captcha {
	c := captchImpl{}
	return &c
}

type captchImpl struct {
	phrase string
	promptMsg string
	errorMsg string
}

func (captcha *captchImpl) ConfirmPhrase(reader io.Reader) error {
	return nil
}

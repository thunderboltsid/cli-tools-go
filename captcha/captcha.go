package captcha

import (
	"io"
)

const (
	defaultCaptchaLength = 6
)

type Captcha interface {
	ConfirmPhrase(reader io.Reader) error
}

func New(opts ...func(*captchaImpl)) Captcha {
	c := captchaImpl{}
	for _, option := range opts {
		option(&c)
	}
	if c.phrase == "" {
		c.phrase = randomString(defaultCaptchaLength)
	}
	return &c
}

type captchaImpl struct {
	// phrase stores the randomly generated string
	phrase string
	// promptMsg stores the message showing prompt asking for confirmation
	promptMsg string
	// errorMsg stores the message showing error in case of failure to confirm phrase
	errorMsg string
}

// Length sets the length of the captcha phrase
func Length(length int) func (captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.phrase = randomString(length)
	}
}

// PromptMessage sets the message showing prompt asking for confirmation
func PromptMessage(msg string) func (captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.promptMsg = msg
	}
}

// ErrorMessage sets the message showing error in case of failure to confirm phrase
func ErrorMessage(msg string) func (captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.errorMsg = msg
	}
}

func (captcha *captchaImpl) ConfirmPhrase(reader io.Reader) error {
	return nil
}

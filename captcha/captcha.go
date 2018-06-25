package captcha

import (
	"io"
	"bufio"
	"strings"
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/thunderboltsid/cli-tools-go/captcha/alphabet"
)

const (
	defaultCaptchaLength = 6
)

var (
	defaultPrintFunc = logrus.Infof
	defaultAlphabet  = alphabet.HollowBlockAlphabet()
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
	if c.print == nil {
		c.print = defaultPrintFunc
	}
	if c.alphabet == nil {
		c.alphabet = defaultAlphabet
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
	// print is the function used for formatted printing
	print func(format string, a ...interface{})
	// alphabet is the alphabet representation set used for rendering the captcha
	alphabet alphabet.Alphabet
}

// Length sets the length of the captcha phrase
func Length(length int) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.phrase = randomString(length)
	}
}

// PromptMessage sets the message showing prompt asking for confirmation
func PromptMessage(msg string) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.promptMsg = msg
	}
}

// ErrorMessage sets the message showing error in case of failure to confirm phrase
func ErrorMessage(msg string) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.errorMsg = msg
	}
}

// PrintFunc sets the print function on the captcha
func PrintFunc(print func(string, ...interface{})) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.print = print
	}
}

// PrintFunc sets the print function on the captcha
func WithAlphabet(a alphabet.Alphabet) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.alphabet = a
	}
}

// ConfirmPhrase prints the prompt message and takes input from reader
func (captcha *captchaImpl) ConfirmPhrase(reader io.Reader) error {
	captcha.print(captcha.promptMsg)
	bufferedReader := bufio.NewReader(reader)
	response, err := bufferedReader.ReadString('\n')
	if err != nil {
		return err
	}
	response = strings.TrimSpace(response)
	if response != captcha.phrase {
		return fmt.Errorf(captcha.errorMsg)
	}
	return nil
}

package captcha

import (
	"io"
	"bufio"
	"strings"
	"github.com/sirupsen/logrus"
	"fmt"
	"github.com/thunderboltsid/cli-tools-go/captcha/alphabet"
	"os"
)

const (
	defaultCaptchaLength = 6
)

var (
	defaultPrintFunc = logrus.Infof
	defaultAlphabet  = alphabet.HollowBlockAlphabet()
)

type Captcha interface {
	ConfirmPhrase() error
}

func defaultCaptcha() *captchaImpl {
	return &captchaImpl{
		phrase:   randomString(defaultCaptchaLength),
		print:    defaultPrintFunc,
		alphabet: defaultAlphabet,
		reader:   os.Stdin,
	}
}

func New(opts ...func(*captchaImpl)) Captcha {
	c := defaultCaptcha()
	for _, option := range opts {
		option(c)
	}
	return c
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
	// reader specifies how the captcha input is read
	reader io.Reader
}

// WithLength sets the length of the captcha phrase
func WithLength(length int) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.phrase = randomString(length)
	}
}

// WithPromptMessage sets the message showing prompt asking for confirmation
func WithPromptMessage(msg string) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.promptMsg = msg
	}
}

// WithErrorMessage sets the message showing error in case of failure to confirm phrase
func WithErrorMessage(msg string) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.errorMsg = msg
	}
}

// WithPrintFunc sets the print function on the captcha
func WithPrintFunc(print func(string, ...interface{})) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.print = print
	}
}

// WithPrintFunc sets the print function on the captcha
func WithAlphabet(alphabet alphabet.Alphabet) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.alphabet = alphabet
	}
}

// WithReader sets the input reader on the captcha
func WithReader(reader io.Reader) func(captcha *captchaImpl) {
	return func(captcha *captchaImpl) {
		captcha.reader = reader
	}
}

// ConfirmPhrase prints the prompt message and takes input from reader
func (captcha *captchaImpl) ConfirmPhrase() error {
	captcha.print(captcha.promptMsg)
	bufferedReader := bufio.NewReader(captcha.reader)
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

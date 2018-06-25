package main

import (
	"github.com/thunderboltsid/cli-tools-go/captcha"
	"log"
)

func main() {
	c, err := captcha.New(captcha.WithPromptMessage("Are you sure you want to make this change?"), captcha.WithErrorMessage("You done fucked up"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err := c.ConfirmPhrase(); err != nil {
		log.Fatalf("Nope: %s", err.Error())
	}
}

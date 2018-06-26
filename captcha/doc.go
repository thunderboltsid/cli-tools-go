/*
Package captcha provides a library to build captcha-style verification mechanisms into your CLI tools. Captchas can
be created using a variety of options passed via the functional options pattern. Each option has a corresponding sane
default so options don't need to be provided but can be for tuning the captcha behaviour.

Example usage (simple): Captcha with default arguments

	c, err := captcha.New()
	if err := c.ConfirmPhrase(); err != nil {
		// handle error
	}

Example usage (complex): Captcha with custom options

	c, err := captcha.New(
		captcha.WithLength(10), // set length of captcha phrase
		captcha.WithPromptMessage("Are you sure you want to make this change?"), // set the prompt message
		captcha.WithErrorMessage("Invalid response to the captcha prompt; aborting change")) // set the error message
	if err := c.ConfirmPhrase(); err != nil {
		// handle error
	}

The resulting output looks something like this:

	 __   __ .  _  . __      __ .  ____ .  ___  .  ___
	 \ \ / / . / | . \ \    / / . |__ / . | _ ) . | _ \
	  \ V /  . | | .  \ \/\/ /  .  |_ \ . | _ \ . |  _/
	   \_/   . |_| .   \_/\_/   . |___/ . |___/ . |_|
	         .     .            .       .       .
	Are you sure you want to make this change?

*/
package captcha

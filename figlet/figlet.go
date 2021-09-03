package figlet

import (
	"fmt"
	"io"
	"strings"

	"github.com/thunderboltsid/cli-tools-go/figlet/alphabet"
)

func defaultFiglet() (*figletImpl, error) {
	alphabet, err := alphabet.New()
	if err != nil {
		return nil, err
	}
	return &figletImpl{
		phrase:   "hello world",
		alphabet: alphabet,
	}, nil
}

// New is a constructor that constructs a figlet string using the functional options pattern
func New(opts ...func(*figletImpl)) error {
	f, err := defaultFiglet()
	if err != nil {
		return fmt.Errorf("unable to create figlet: %s", err.Error())
	}
	for _, option := range opts {
		option(f)
	}

	renderString(f.phrase, f.alphabet, f.writer)

	return nil
}

type figletImpl struct {
	// phrase stores the message string
	phrase string
	writer io.Writer
	// alphabet is the alphabet representation set used for rendering the figlet
	alphabet alphabet.Alphabet
}

// WithWriter sets the writer
func WithWriter(writer io.Writer) func(figlet *figletImpl) {
	return func(figlet *figletImpl) {
		figlet.writer = writer
	}
}

// WithMsg sets the message
func WithMsg(msg string) func(figlet *figletImpl) {
	return func(figlet *figletImpl) {
		figlet.phrase = strings.ToUpper(msg)
	}
}

// WithAlphabet sets the alphabet used for rendering the captcha
func WithAlphabet(alphabet alphabet.Alphabet) func(figlet *figletImpl) {
	return func(figlet *figletImpl) {
		figlet.alphabet = alphabet
	}
}

// renderString renders a given string into corresponding alphabet representation and prints it using the provided
// print method
func renderString(str string, alphabet alphabet.Alphabet, writer io.Writer) {
	var out []string
	for i := 0; i < len(str); i++ {
		char := alphabet.RenderMap()[str[i:i+1]]
		if out == nil {
			out = make([]string, len(char))
		}
		for k, v := range char {
			if out[k] == "" {
				out[k] = v
			} else {
				out[k] = fmt.Sprintf("%s.%s", out[k], v)
			}
		}
	}
	for _, v := range out {
		fprintln(writer, "%s", v)
	}
}

func fprintln(writer io.Writer, format string, a ...interface{}) {
	_, err := fmt.Fprintf(writer, format, a...)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Fprint(writer, "\n")
	if err != nil {
		fmt.Println(err)
	}
}

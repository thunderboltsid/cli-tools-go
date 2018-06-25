package captcha

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thunderboltsid/cli-tools-go/captcha/alphabet"
)

func Test__New_creates_new_captcha(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	require.NotNil(t, c)
}

func Test__New_creates_captcha_with_default_length(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.Len(t, ci.phrase, defaultCaptchaLength)
}

func Test__New_creates_captcha_with_default_print_function(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.NotNil(t, ci.print)
	assert.Equal(t, reflect.ValueOf(ci.print).Pointer(), reflect.ValueOf(defaultPrintFunc).Pointer())
}

func Test__New_creates_captcha_with_default_alphabet(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.NotNil(t, ci.alphabet)
	alphabet, err := alphabet.New()
	require.NoError(t, err)
	assert.Equal(t, alphabet, ci.alphabet)
}

func Test__New_creates_captcha_with_default_reader(t *testing.T) {
	c, err := New()
	require.NoError(t, err)
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.NotNil(t, ci.reader)
	assert.Equal(t, os.Stdin, ci.reader)
}

func Test__New_creates_captcha_with_specified_length(t *testing.T) {
	c, err := New(WithLength(10))
	require.NoError(t, err)
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.Len(t, ci.phrase, 10)
}

func Test__New_creates_captcha_with_specified_prompt_and_error_msg(t *testing.T) {
	c, err := New(WithPromptMessage("prompt"), WithErrorMessage("error"))
	require.NoError(t, err)
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.Equal(t, "prompt", ci.promptMsg)
	assert.Equal(t, "error", ci.errorMsg)
}

func Test__New_creates_captcha_with_specified_print_function(t *testing.T) {
	c, err := New(WithPrintFunc(log.Fatalf))
	require.NoError(t, err)
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.NotNil(t, ci.print)
	assert.Equal(t, reflect.ValueOf(ci.print).Pointer(), reflect.ValueOf(log.Fatalf).Pointer())
}

package captcha

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"reflect"
	"log"
	"github.com/thunderboltsid/cli-tools-go/captcha/alphabet"
)

func Test__New_creates_new_captcha(t *testing.T) {
	c := New()
	require.NotNil(t, c)
}

func Test__New_creates_captcha_with_default_length(t *testing.T) {
	c := New()
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.Len(t, ci.phrase, defaultCaptchaLength)
}

func Test__New_creates_captcha_with_default_print_function(t *testing.T) {
	c := New()
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.NotNil(t, ci.print)
	assert.Equal(t, reflect.ValueOf(ci.print).Pointer(), reflect.ValueOf(defaultPrintFunc).Pointer())
}

func Test__New_creates_captcha_with_default_alphabet(t *testing.T) {
	c := New()
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.NotNil(t, ci.alphabet)
	assert.Equal(t, alphabet.HollowBlockAlphabet(), ci.alphabet)
}

func Test__New_creates_captcha_with_specified_length(t *testing.T) {
	c := New(WithLength(10))
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.Len(t, ci.phrase, 10)
}

func Test__New_creates_captcha_with_specified_prompt_and_error_msg(t *testing.T) {
	c := New(WithPromptMessage("prompt"), WithErrorMessage("error"))
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.Equal(t, "prompt", ci.promptMsg)
	assert.Equal(t, "error", ci.errorMsg)
}

func Test__New_creates_captcha_with_specified_print_function(t *testing.T) {
	c := New(WithPrintFunc(log.Fatalf))
	require.NotNil(t, c)
	ci, ok := c.(*captchaImpl)
	require.True(t, ok)
	assert.NotNil(t, ci.print)
	assert.Equal(t, reflect.ValueOf(ci.print).Pointer(), reflect.ValueOf(log.Fatalf).Pointer())
}
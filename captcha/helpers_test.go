package captcha

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_randomString_has_correct_length(t *testing.T) {
	length := 10
	randomStr := randomString(length)
	require.Len(t, randomStr, length)
}

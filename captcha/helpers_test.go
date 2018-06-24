package captcha

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_helpers_randomString(t *testing.T) {
	length := 10
	randomStr := randomString(length)
	require.Len(t, randomStr, length)
}

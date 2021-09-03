package figlet

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(WithMsg("yolo"), WithWriter(os.Stdout))
	assert.NoError(t, err)
}
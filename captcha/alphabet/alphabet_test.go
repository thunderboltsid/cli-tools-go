package alphabet

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test__New_creates_alphabet_with_hollowBlockAlphabet_by_default(t *testing.T) {
	alphabet, err := New()
	require.NoError(t, err)
	require.NotNil(t, alphabet)
	ai, ok := alphabet.(*alphabetImpl)
	require.True(t, ok)
	assert.Equal(t, hollowBlockAlphabet, string(ai.alphabet))
}

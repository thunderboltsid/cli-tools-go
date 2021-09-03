package alphabet

import (
	"encoding/json"
)

// Alphabet interface defines the contract that a given alphabet needs to fulfil
type Alphabet interface {
	RenderMap() map[string][]string
}

// New is a constructor for Alphabet objects using functional options pattern
func New(opts ...func(impl *alphabetImpl)) (Alphabet, error) {
	alphabet := defaultAlphabet()
	for _, option := range opts {
		option(alphabet)
	}
	err := json.Unmarshal(alphabet.alphabet, &alphabet.characterMap)
	if err != nil {
		return nil, err
	}
	return alphabet, nil
}

func defaultAlphabet() *alphabetImpl {
	return &alphabetImpl{
		alphabet: []byte(HollowBlockAlphabet),
	}
}

// WithAlphabet specifies the alphabet e.g. HollowBlockAlphabet
func WithAlphabet(alphabet string) func(alphabet *alphabetImpl) {
	return func(a *alphabetImpl) {
		a.alphabet = []byte(alphabet)
	}
}

type alphabetImpl struct {
	// take a look at HollowBlockAlphabet to understand how alphabet looks like
	alphabet     []byte
	characterMap map[string][]string
}

// RenderMap returns the map which defines how a given character corresponds to it's ASCII art equivalent
func (alphabet *alphabetImpl) RenderMap() map[string][]string {
	return alphabet.characterMap
}

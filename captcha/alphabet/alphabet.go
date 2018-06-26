package alphabet

import (
	"encoding/json"
)

type Alphabet interface {
	RenderMap() map[string][]string
}

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

func (alphabet *alphabetImpl) RenderMap() map[string][]string {
	return alphabet.characterMap
}

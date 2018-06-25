package alphabet

import (
	"encoding/json"
)

type Alphabet interface {
	RenderMap() map[string][]string
}

type alphabetImpl struct {
	characterMap map[string][]string
}

func (alphabet *alphabetImpl) RenderMap() map[string][]string {
	return alphabet.characterMap
}

func HollowBlockAlphabet() Alphabet {
	alphabet := alphabetImpl{
		characterMap: make(map[string][]string),
	}
	json.Unmarshal([]byte(hollowBlockCharacterMap), &alphabet.characterMap)
	return &alphabet
}

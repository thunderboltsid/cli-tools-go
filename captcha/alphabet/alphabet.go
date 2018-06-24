package alphabet

import (
	"encoding/json"
)

type Alphabet interface {
	RenderMap() map[string]string
}

type alphabetImpl struct {
	characterMap map[string]string
}

func (alphabet *alphabetImpl) RenderMap() map[string]string {
	return alphabet.characterMap
}

func HollowBlockAlphabet() (Alphabet, error) {
	alphabet := alphabetImpl{
		characterMap:make(map[string]string),
	}
	if err := json.Unmarshal([]byte(hollowBlockCharacterMap), alphabet.characterMap); err != nil {
		return nil, err
	}
	return &alphabet, nil
}

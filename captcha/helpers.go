package captcha

import (
	"math/rand"
	"time"
	"github.com/thunderboltsid/cli-tools-go/captcha/alphabet"
	"fmt"
)

// letterBytes represents the characters to be used (several omitted for readability) while building a random string for captcha
const letterBytes = "123479ABCDEFGHIJKLMNPRTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randomString(length int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func renderCaptcha(str string, alphabet alphabet.Alphabet) []string {
	var out []string
	for i := 0; i < len(str); i++ {
		char := alphabet.RenderMap()[str[i:i+1]]
		if out == nil {
			out = make([]string, len(char))
		}
		for k, v := range char {
			if out[k] == "" {
				out[k] = v
			} else {
				out[k] = fmt.Sprintf("%s.%s", out[k], v)
			}
		}
	}
	return out
}

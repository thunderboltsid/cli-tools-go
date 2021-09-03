package captcha

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/thunderboltsid/cli-tools-go/figlet/alphabet"
)

// letterBytes represents the characters to be used (several omitted for readability) while building a random string for captcha
const letterBytes = "123479ABCDEFGHIJKLMNPRTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// randomString generates a random string of a given length
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

// renderString renders a given string into corresponding alphabet representation and prints it using the provided
// print method
func renderString(str string, alphabet alphabet.Alphabet, print func(string, ...interface{})) {
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
	for _, v := range out {
		print("%s", v)
	}
}

func println(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Print("\n")
}

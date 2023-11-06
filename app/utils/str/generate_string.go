package str

import (
	"math/rand"
	"time"
)

func GenerateString(length int, chartype string) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	switch chartype {
	case "char":
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		break
	case "digit":
		charset = "0123456789"
		break
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	b := make([]byte, length)

	for i := range b {
		b[i] = charset[r1.Intn(len(charset))]
	}

	return string(b)
}

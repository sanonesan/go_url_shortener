package services

import (
	"crypto/sha256"
	"strings"
)

const (
	forbidden = 63
	decision  = 45
	alphabet  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

// generateShortURL
func GenerateShortURL(long string) string {
	hash := sha256.Sum224([]byte(long))
	res := strings.Builder{}
	j := 0
	for i := 0; i < 10; i++ {
		curr := hash[i] & ((1 << 6) - 1)
		jOld := j
		for curr == forbidden {
			curr ^= hash[j+10] & ((1 << 6) - 1)
			j = (j + 1) % 18
			if j == jOld &&
				curr == forbidden {
				curr ^= decision
			}
		}
		symbol := alphabet[curr]
		res.WriteByte(symbol)
	}
	return res.String()
}

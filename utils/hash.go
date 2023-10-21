package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func getSHA256Binary(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}

func CreateHashFromString(s string) string {
	b := getSHA256Binary(s)
	h := hex.EncodeToString(b)
	return h
}

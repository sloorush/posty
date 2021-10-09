package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func BytesToString(data []byte) string {
	return hex.EncodeToString(data)
}

func Hash(word string) (ret string) {
	data := []byte(word)
	hash := NewSHA256(data)

	// fmt.Println(hash)

	ret = BytesToString(hash)
	// fmt.Println(ret)
	return ret
}

func VerifyHash(word string, checkHash string) (ret bool) {
	newHash := Hash(word)

	return newHash == checkHash
}

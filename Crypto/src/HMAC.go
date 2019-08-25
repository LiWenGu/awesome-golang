package src

import (
	"crypto/hmac"
	"crypto/sha256"
)

func GenerateHMAC(src, key []byte) []byte {
	hasher := hmac.New(sha256.New, key)

	mac := hasher.Sum(nil)
	return mac
}

func VerifyHMAC(src, key, mac1 []byte) bool {
	mac2 := GenerateHMAC(src, key)
	return hmac.Equal(mac1, mac2)
}

package src

import (
	"fmt"
	"testing"
)

func TestGenerateHMAC(t *testing.T) {
	src := []byte("我是一个明文的hash值和公钥")
	key := []byte("12345678")
	mac1 := GenerateHMAC(src, key)
	fmt.Printf("mac1: %x\n", mac1)
	isEqual := VerifyHMAC(src, key, mac1)
	fmt.Println(isEqual)
}

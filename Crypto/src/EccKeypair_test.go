package src

import (
	"fmt"
	"testing"
)

func TestGenerateEccKeypair(t *testing.T) {
	generateEccKeypair()

	src := []byte("GOlang, 不支持 ECC 加解密，支持 ECC 签名")
	signature := EccSignData(ECCPrivateKeyFile, src)
	fmt.Println(signature)
	isVerify := EccVerifySign(ECCPublicKeyFile, src, signature)
	fmt.Println(isVerify)
}

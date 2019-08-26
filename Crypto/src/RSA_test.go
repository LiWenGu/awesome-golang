package src

import (
	"fmt"
	"testing"
)

func TestRsaGenKey(t *testing.T) {
	RsaGenKey(4096)

	src := []byte("我是一个明文")
	fmt.Println("加密前：" + string(src))
	data := RSAPublicEncrypto(src, PublicKeyFile)
	fmt.Println("加密后：" + string(data))
	data = RSAPrivateDecrypt(data, PrivateKeyFile)
	fmt.Println("解密后：" + string(data))
}

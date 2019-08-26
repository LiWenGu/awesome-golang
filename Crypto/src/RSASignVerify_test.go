package src

import (
	"fmt"
	"testing"
)

// 私钥签名
func TestRsaSign(t *testing.T) {
	src := []byte("我是一个明文哦")
	signature := RsaSignData(PrivateKeyFile, src)
	fmt.Println("密文：" + string(signature))
	isVerify := RsaVerifySignature(signature, src, PublicKeyFile)
	if isVerify {
		fmt.Println("验证成功")
	} else {
		fmt.Println("验证失败")
	}
}

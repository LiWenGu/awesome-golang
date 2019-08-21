package src

import (
	"fmt"
	"testing"
)

func TestDES(t *testing.T) {
	src := []byte("我是一个不知道有多长的明文")
	key := []byte("12345678")
	str := encryptoDES(src, key)
	fmt.Println("加密之后的密文：" + string(str))
	str = decryptoDES(str, key)
	fmt.Println("解密之后的明文：" + string(str))
}

func Test3DES(t *testing.T) {
	src := []byte("我是一个不知道有多长的明文")
	key := []byte("12345678abcdefgh87654321")
	str := encrypto3DES(src, key)
	fmt.Println("加密之后的密文：" + string(str))
	str = decrypto3DES(str, key)
	fmt.Println("解密之后的明文：" + string(str))
}

func TestAES(t *testing.T) {
	src := []byte("我是一个不知道有多长的明文")
	// TODO 这里会报错，因为秘钥长度是 17 byte，去掉一个字符就好了
	key := []byte("1234abcd4321dcbaw")
	str := encryptoAES(src, key)
	fmt.Println("加密之后的密文：" + string(str))
	str = decryptoAES(str, key)
	fmt.Println("解密之后的明文：" + string(str))
}

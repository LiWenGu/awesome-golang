package src

import (
	"fmt"
	"testing"
)

func TestDES(t *testing.T) {
	fmt.Println("===== des 加解密 ====")
	src := []byte("少壮不努力，老大徒伤悲")
	key := []byte("12345678")
	str := encryptoDES(src, key)
	fmt.Println("加密之后的密文：" + string(str))
	str = decryptoDES(str, key)
	fmt.Println("解密之后的明文：" + string(str))
}

func Test3DES(t *testing.T) {
	fmt.Println("===== 3des 加解密 ====")
	src := []byte("少壮不努力，老大徒伤悲")
	key := []byte("12345678abcdefgh87654321")
	str := encrypto3DES(src, key)
	fmt.Println("加密之后的密文：" + string(str))
	str = decrypto3DES(str, key)
	fmt.Println("解密之后的明文：" + string(str))
}

func TestAES(t *testing.T) {
	fmt.Println("===== aes 加解密 ====")
	src := []byte("少壮不努力，老大徒伤悲")
	key := []byte("1234abcd4321dcbaw")
	str := encryptoAES(src, key)
	fmt.Println("加密之后的密文：" + string(str))
	str = decryptoAES(str, key)
	fmt.Println("解密之后的明文：" + string(str))
}

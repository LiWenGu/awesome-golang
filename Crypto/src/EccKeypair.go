package src

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

// 生成 ecc 秘钥对
func generateEccKeypair() {
	// 1. 选择一个椭圆曲线
	curve := elliptic.P256()
	// 2. 生成秘钥对
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr("generate key failed!", err)
	// 3. 将 ecdsa 私钥序列化为 ASN.1 DER 编码
	derText, err := x509.MarshalECPrivateKey(privateKey)
	block := pem.Block{
		Type:    "ECC PRIVATE KEY",
		Headers: nil,
		Bytes:   derText,
	}
	fileHandler, err := os.Create(ECCPrivateKeyFile)
	checkErr("os.create failed", err)
	defer fileHandler.Close()

	err = pem.Encode(fileHandler, &block)
	checkErr("pem decode failed", err)

	// 生成公钥
	publicKey := privateKey.PublicKey
	// 通用的序列化方式
	derText, err = x509.MarshalPKIXPublicKey(&publicKey)
	checkErr("MarshalPKIXPublicKey", err)
	block = pem.Block{
		Type:    "ECC PUBLICK KEY",
		Headers: nil,
		Bytes:   derText,
	}
	fileHandler, err = os.Create(ECCPublicKeyFile)
	checkErr("os.Create", err)
	defer fileHandler.Close()
	err = pem.Encode(fileHandler, &block)
	checkErr("pem encode failed", err)
}

type Signature struct {
	r *big.Int
	s *big.Int
}

// 私钥签名
func EccSignData(filename string, src []byte) Signature {
	// 1. 读取私钥，解码
	info, err := ioutil.ReadFile(filename)
	checkErr("ioutil.ReaderFile failed", err)
	block, _ := pem.Decode(info)
	derText := block.Bytes
	privateKey, err := x509.ParseECPrivateKey(derText)
	checkErr("x509.ParseECPrivateKey failed", err)
	// 2. 对原文生成哈希
	hash := sha256.Sum256(src)
	// 3. 使用私钥签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	checkErr("sign failed", err)
	return Signature{r, s}
}

// PKCS#1定义了RSA公钥函数的基本格式标准，特别是数字签名
// IETF的安全领域的公钥基础实施（PKIX）工作组正在为互联网上使用的公钥证书定义一系列的标准。PKIX工作组在1995年10月成立。
func EccVerifySign(filename string, src []byte, sign Signature) bool {
	// 1. 读取公钥、解码
	info, err := ioutil.ReadFile(filename)
	checkErr("ioutil.ReaderFile failed", err)
	block, _ := pem.Decode(info)
	derText := block.Bytes
	publicKeyInterface, err := x509.ParsePKIXPublicKey(derText)
	checkErr("x509.ParsePKCS1PublicKey", err)
	publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("断言失败，非ecsa公钥！")
	}
	// 2. 对原文生成哈希
	hash := sha256.Sum256(src)
	// 3. 校验，使用公钥验证 hash 值和两个大整数 r、s 构成的签名，并返回签名是否合法
	return ecdsa.Verify(publicKey, hash[:], sign.r, sign.s)
}

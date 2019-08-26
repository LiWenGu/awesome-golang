package src

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

// 私钥签名
func RsaSignData(filename string, src []byte) []byte {
	info, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(info)
	derText := block.Bytes
	privateKey, err := x509.ParsePKCS1PrivateKey(derText)
	if err != nil {
		panic(nil)
	}
	// 1. 获取原文的 hash 值
	hash := sha256.Sum256(src)
	// 2. 对原文 hash 值使用私钥进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		panic(err)
	}
	return signature
}

// 公钥认证
func RsaVerifySignature(sign []byte, src []byte, filename string) bool {
	info, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(info)
	derText := block.Bytes
	publicKey, err := x509.ParsePKCS1PublicKey(derText)
	if err != nil {
		panic(nil)
	}
	// 1. 获取原文的 hash 值
	hash := sha256.Sum256(src)
	// 2. 公钥解密签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], sign)
	if err != nil {
		return false
	}
	return true
}

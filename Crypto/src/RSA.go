package src

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func RsaGenKey(bits int) {
	// 1. 使用 rsa 的 GenerateKey 方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	// 2. 通过 x509 标准将得到的 rsa 私钥序列化为 ASN.1 的 DER 编码字符串
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. 将私钥字符串设置到 pem 格式块中
	block := pem.Block{
		Type:    "RSA Private Key",
		Headers: nil,
		Bytes:   privateStream,
	}
	// 4. 通过 pem 将设置好的数据进行编码，写入文件
	privateFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(privateFile, &block)
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()

	// 1. 从得到的私钥对象中将公钥信息取出
	pubKey := privateKey.PublicKey
	// 2. 通过 x509 标准将得到的 rsa 公钥序列化为字符串
	pubStream := x509.MarshalPKCS1PublicKey(&pubKey)
	// 3. 将公钥字符串设置到 pem 格式块中
	block = pem.Block{
		Type:    "RSA Public key",
		Headers: nil,
		Bytes:   pubStream,
	}
	// 4. 通过 pem 将设置好的数据进行编码，写入文件
	pubFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(pubFile, &block)
	if err != nil {
		panic(err)
	}
	defer pubFile.Close()
}

// 公钥加密函数
func RSAPublicEncrypto(src []byte, pathName string) []byte {
	// 1. 得到公钥文件中的公钥
	file, err := os.Open(pathName)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	recevBuf := make([]byte, info.Size())
	file.Read(recevBuf)
	// 2. 将得到的字符串解码
	block, _ := pem.Decode(recevBuf)
	// 3. 使用 xx509 将编码之后公钥解析出来
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(nil)
	}
	// 4. 使用公钥加密
	msg, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	if err != nil {
		panic(err)
	}
	return msg
}

// 使用私钥解密
func RSAPrivateDecrypt(src []byte, pathName string) []byte {
	// 1. 打开私钥文件
	file, err := os.Open(pathName)
	if err != nil {
		panic(err)
	}
	// 2. 读私钥文件内容
	info, _ := file.Stat()
	recevBuf := make([]byte, info.Size())
	file.Read(recevBuf)
	// 3. 将得到的字符串解码
	block, _ := pem.Decode(recevBuf)
	// 4. 通过 x509 还原私钥数据
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	msg, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
	if err != nil {
		panic(err)
	}
	return msg
}

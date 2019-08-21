package src

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

// 填充最后一个分组的函数
// src - 原始函数
// blockSize - 每个分组的数据长度
func paddingText(src []byte, blockSize int) []byte {
	// 1. 求出最后一个分组要填充多少个字节
	padding := blockSize - len(src)%blockSize
	// 2. 创建新的切片，切片的字节数为 padding，并初始化，每个字节的值为 padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 3. 将创建出的新切片和原始数据进行连接
	newText := append(src, padText...)
	// 4. 返回新的字符串
	return newText
}

// 删除末尾填充的字节
func unPaddingText(src []byte) []byte {
	// 1. 求出要处理的切片的长度
	len := len(src)
	// 2. 取出最后一个字符
	number := int(src[len-1])
	// 3. 将切片末尾的number个字节删除
	newText := src[:len-number]
	return newText
}

// 使用 des 加密
// src - 明文
// key - 秘钥
func encryptoDES(src []byte, key []byte) []byte {
	// 1. 创建并返回一个使用 DES 算法的 cipher.Block 接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 对最后一个明文分组填充
	src = paddingText(src, block.BlockSize())
	// 3. 创建一个密码分组为链接模式，底层使用 DES 加密的 BlockMode 接口
	// 初始化向量，必须长度和秘钥长度相同
	iv := []byte("aaaabbbb")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4. 加密连续的数据块
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// 使用 des 解密
func decryptoDES(src []byte, key []byte) []byte {
	// 1. 创建并返回一个使用 DES 算法的 cipher.Block 接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个密码分组为链接模式的，底层使用 DES 解密的 BlockMode 接口
	iv := []byte("aaaabbbb")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3. 数据块解密
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	// 4. 去掉最后一组的填充数据
	newText := unPaddingText(dst)
	return newText
}

func encrypto3DES(src, key []byte) []byte {
	// 1. 创建并返回一个使用 3DES 算法的 cipher.Block 接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	src = paddingText(src, block.BlockSize())
	// 初始化向量，必须长度和秘钥长度相同
	iv := []byte("aaaabbbb")
	// 2. 创建一个密码分组为链接模式的，底层使用 3DES 加密的 BlockMode 接口
	blockMode := cipher.NewCBCEncrypter(block, iv)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

func decrypto3DES(src, key []byte) []byte {
	// 1. 创建并返回一个使用 3DES 算法的 cipher.Block 接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	// 初始化向量，必须长度和秘钥长度相同
	iv := []byte("aaaabbbb")
	// 2. 创建一个密码分组为链接模式的，底层使用 3DES 解密的 BlockMode 接口
	blockMode := cipher.NewCBCDecrypter(block, iv)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	newText := unPaddingText(dst)
	return newText
}

func encryptoAES(src, key []byte) []byte {
	// 1. 创建并返回一个使用 AES 算法的 cipher.Block 接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	src = paddingText(src, block.BlockSize())
	// 初始化向量，必须长度和秘钥长度相同
	iv := []byte("aaaabbbbccccdddd")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

func decryptoAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 初始化向量，必须长度和秘钥长度相同
	iv := []byte("aaaabbbbccccdddd")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	newText := unPaddingText(dst)
	return newText
}

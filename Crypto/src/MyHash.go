package src

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Str_1(src []byte) string {
	res := md5.Sum(src)
	myres := hex.EncodeToString(res[:])
	return myres
}

func GetMD5Str_2(src []byte) string {
	hash := md5.New()
	// io.WriteString(hash, string(src))
	hash.Write(src)
	res := hash.Sum(nil)
	// 散列值格式化
	myres := hex.EncodeToString(res)
	return myres
}

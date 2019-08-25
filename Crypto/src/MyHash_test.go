package src

import (
	"fmt"
	"testing"
)

func TestGetMD5Str_1(t *testing.T) {
	str := GetMD5Str_1([]byte("我是一个明文"))
	// e8e0915549a588dca746a1db55c75b35
	fmt.Println(str)
}

func TestGetMD5Str_2(t *testing.T) {
	str := GetMD5Str_2([]byte("我是一个明文"))
	// e8e0915549a588dca746a1db55c75b35
	fmt.Println(str)
}

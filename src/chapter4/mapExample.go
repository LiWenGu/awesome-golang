package main

import "fmt"

func main() {
	// 创建空映射
	dict1 := make(map[string]int)
	dict2 := map[string]string{"red": "#da1337"}
	// 创建 nil 的映射
	var dict3 map[string]string
	// 报错！修改映射
	dict3["x"] = "x"

	// 获取映射
	value, exists := dict2["red"]
	if exists {
		fmt.Println(value)
	}
	valuez := dict2["red"]
	if value != "" {
		fmt.Printf(valuez)
	}

	// 遍历映射
	for key, value := range dict2 {
		fmt.Println(key, value)
	}
	// 删除映射
	delete(dict2, "red")
}

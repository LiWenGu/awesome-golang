package main

import "fmt"

func main() {
	// 创建长度和容量都是5的切片
	slice1 := make([]string, 5)
	slice1[0] = "1"
	// 创建长度是3 容量是5 的切片
	slice0 := make([]string, 3, 5)
	slice0[0] = "1"
	// 创建长度和容量都是3的切片
	slice2 := []int{10, 20, 30}
	slice2[1] = 1
	// 创建数组
	array := [3]int{10, 20, 30}
	array[1] = 0
	// 使用切片

	slicez := []int{10, 20, 30, 40}
	// 长度：3-1=2，容量：5-1=4
	slicex := slicez[1:3]
	slicex[0] = 1
	// 扩容，1000以内 2.0，1000+ 1.25
	slices := append(slicez, 50, 50, 50)
	fmt.Printf("%v\n", slices)

	for index, value := range slices {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// len() 切片的长度， cap() 切片的容量
	for index := 2; index < len(slices); index++ {
		fmt.Printf("%s", slices[index])
	}
}

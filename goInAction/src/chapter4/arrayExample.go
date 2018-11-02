package main

func main() {
	array := [5]int{10}
	array[2] = 2

	// 指向 int 的指针数组
	arrayz := [5]*int{0: new(int), 1: new(int)}
	*arrayz[0] = 10

	// 二维数组
	var array2 [4][2]int
	array2[0][0] = 1

	array22 := [4][2]int{0: {0: 1, 1: 2}}
	array22[0][0] = 1

	// 分配 8MB 的数组
	var arrayBig [1e6]int
	// 大数组使用指针更省内存
	foo(&arrayBig)

}

func foo(array *[1e6]int) {

}

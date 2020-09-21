package main

import (
	"fmt"
)

// 声明一个函数类型
type cb func(int) int

func test1(x int) int {
	return x
}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调函数，x：%d\n", x)
	return x
}
func cat(x int, f cb) int {
	return f(x)
}

func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调1221，x：%d\n", x)
		return x
	})
	fmt.Println(cat(20, test1))

}

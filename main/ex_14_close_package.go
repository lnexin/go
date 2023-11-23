package main

import "fmt"

func main() {
	nextInt := intSeq()
	// inntSeq 函数返回两一个在intSeq 函数体内定义的匿名函数
	// 这个返回的函数使用闭包的方式  隐藏变量 i

	// 我们调用 intSeq 函数，将返回值（也是一个函数）赋给nextInt, 这个函数的值包含了自己的值i ,这样在每次
	// 调用nextInt 时都会更新i的值

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

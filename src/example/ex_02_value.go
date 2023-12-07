package main

import "fmt"

func main() {
	fmt.Println("hello world" + "haha")

	fmt.Println("1+1=", 1+1)

	fmt.Println(7 / 3)

	fmt.Println(7.0 / 3.0)

	fmt.Println(true && false)

	for i := 0; i < 100; i++ {
		x := i
		x++
		fmt.Println(x)
		i++
	}
}

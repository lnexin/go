package main

import "fmt"

func main() {
	i := plus(12, 34)
	fmt.Println(i)

	i2, i3 := vals(4)
	fmt.Println(i2)
	fmt.Println(i3)
}

func plus(a int, b int) int {
	return a + b
}

func vals(a int) (int, int) {
	return a, a - 1
}

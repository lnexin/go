package main

import "fmt"

func main() {
	t1 := t1()
	fmt.Println(t1)

	t2 := t2()
	fmt.Println(t2)
}

func t1() int {
	a := 1
	defer func() {
		a++
	}()
	return a
}

func t2() (a int) {
	defer func() {
		a = a + 3
	}()
	return 3
}

// func main() {
// 	x := 1
// 	y := 2
// 	defer calc("A", x, calc("B", x, y))

// 	x = 3
// 	defer calc("C", x, calc("D", x, y))
// }

// func calc(s string, x, y int) int {
// 	ret := x + y
// 	fmt.Println(s, x, y, ret)
// 	return ret
// }

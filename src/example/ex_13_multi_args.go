package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))

	fmt.Println(sum())

	sls := []int{1, 2, 3, 4}
	fmt.Println(sum(sls...))
}

func sum(nums ...int) int {
	s := 0
	for _, i := range nums {
		s += i
	}
	return s
}

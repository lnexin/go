package main

import (
	"fmt"
)

func main() {
	var sli_1 []int
	print(1, sli_1)

	var sli_2 = []int{}
	print(2, sli_2)

	var sli_3 = []int{1, 2, 3, 4, 5}
	print(3, sli_3)

	sli_4 := []int{1, 2, 3, 7}
	print(4, sli_4)

	var sli_5 []int = make([]int, 5, 8)
	print(5, sli_5)

	sli_6 := make([]int, 5, 9)
	print(6, sli_6)
}

func print(inx int, sli []int) {
	fmt.Printf("inx:%d len:%d, cap:%d, slice: %v\n", inx, len(sli), cap(sli), sli)
}

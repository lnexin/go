package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println(len(s))
	for i := 0; i < 3; i++ {
		s[i] = ("-" + strconv.Itoa(i))
		fmt.Println(len(s))
	}
	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, "duxib", "xib")
	fmt.Println(s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := c[2:4]
	fmt.Println(l)

	l = c[:3]
	fmt.Println(l)

	l = c[3:]
	fmt.Println(l)

	martix := make([][]int, 3)
	for i := 0; i < len(martix); i++ {
		innerLen := i + 1
		martix[i] = make([]int, innerLen)
		for j := 0; j < len(martix[i]); j++ {
			martix[i][j] = i + j
		}
	}
	fmt.Println(martix)
}

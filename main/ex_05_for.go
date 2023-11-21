package main

import "fmt"

func main() {
	i := 13
	for i < 23 {
		fmt.Println(i)
		i++
	}

	for i = 0; i < 13; i++ {
		fmt.Println("i=", i)
	}
	i = 10
	for {
		if i > 15 {
			break
		}
		fmt.Println("lool->", i)
		i++
	}
}

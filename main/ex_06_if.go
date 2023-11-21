package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("s")
	} else {
		fmt.Println("22")
	}

	if n := 9; n < 0 {
		fmt.Println("333")
	} else if n < 10 {
		fmt.Println("444")
	} else {
		fmt.Println("has")
	}

}

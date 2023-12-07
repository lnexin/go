package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println(i)

	switch i {
	case 2:
		fmt.Println("---")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Wednesday, time.Thursday:
		fmt.Println("3,4")
	case time.Tuesday:
		fmt.Println("2")

	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("after noon")
	default:
		fmt.Println("before noon")

	}
}

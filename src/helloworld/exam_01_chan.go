package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	// go func() {
	// 	fmt.Println("go routine")
	// }()

	ch1 := make(chan string, 2)

	ch1 <- "a"
	ch1 <- "bbc"
	go func() {
		for {

			c := <-ch1
			fmt.Println(c)
		}

	}()

	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

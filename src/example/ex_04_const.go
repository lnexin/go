package main

import (
	"fmt"
	"math"
)

const s string = "final static variable"

func main() {
	fmt.Println(s)

	const n = 50000000
	const d = n / 3
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}

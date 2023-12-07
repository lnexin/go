package main

import "fmt"

func main() {
	var p1 = make(map[int]string)
	// p1 =
	// 	1: "ss",
	// 	2:"ff",
	// 	3:"vv"

	// }
	p1[1] = "das"
	// p1[2] = "a"

	fmt.Println(p1)

	var p2 = map[int]string{}

	p2[2] = "Dc"
	fmt.Println(p2)

	p3 := map[int]string{
		4: "$dd",
		5: "%asd",
		6: "^qwe",
	}
	fmt.Println(p3)
}

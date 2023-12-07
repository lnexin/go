package main

import "fmt"

func main() {

	// person := []string{"Tom", "Jack", "James"}
	person := map[int]string{
		1: "T",
		2: "Amone",
		3: "Jax",
	}
	// fmt.Printf("len: %d, cap: %d, array=%v\n", len(person), cap(person), person)
	fmt.Printf("len: %d,  array=%v\n", len(person), person)
	fmt.Println("-----------------")

	for k, v := range person {
		fmt.Printf("person [%d]: %s\n", k, v)
	}

	fmt.Println("-----------------")

	for i := range person {
		fmt.Printf("person[%d]:%s\n", i, person[i])
	}
	fmt.Println("-----------------")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]:%s\n", i, person[i])
	}
	fmt.Println("-----------------")

	for _, name := range person {
		fmt.Println("name:", name)
	}
}

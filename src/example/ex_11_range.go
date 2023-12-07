package main

import "fmt"

func main() {
	nums := []int{11, 12, 13}
	sum := 0
	for _, num := range nums {
		//fmt.Println(i)

		sum += num
		fmt.Println(sum)
	}
	fmt.Println(sum)

	for i, _ := range nums {
		if i == 2 {
			fmt.Println(i)
		}
	}

	kv := map[string]string{"a": "1", "b": "2"}
	for k, v := range kv {
		fmt.Println(k, "-", v)
	}

	for i, c := range "golang" {
		fmt.Println(i, " -> ", c)
	}
}

package main

import "fmt"

func main() {
	sli := []int{4, 5, 6}
	print(1, sli)

	sli = append(sli, 7)

	print(2, sli)

	sli = append(sli, 13, 14, 15)
	print(3, sli)

	sli = append(sli[:2], sli[2+3:]...)
	print(-1, sli)
}

func print(inx int, sli []int) {
	fmt.Printf("inx:%d len:%d, cap:%d, slice: %v\n", inx, len(sli), cap(sli), sli)
}

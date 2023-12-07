package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["du"] = "duxib"
	m["lhc"] = "lihaochen"

	fmt.Println(m)

	fmt.Println(m["lhc"])

	fmt.Println(len(m))

	delete(m, "lhc")
	fmt.Println(m)

	_, prs := m["haha"]
	fmt.Println(prs)

	n := map[string]int{"dx": 1, "l": 2}
	fmt.Println(n)
}

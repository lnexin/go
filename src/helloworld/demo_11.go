package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "Jacl"
	p1.Age = 23

	fmt.Println(p1)

	// var p2 = Person{Name: "sss", age: 11}
	// fmt.Println(p2)

	// p3 := Person{Name: "TOme2", age: 2222}
	// fmt.Println(p3)

	// p4 := struct {
	// 	statusx int
	// 	agex    int
	// }{statusx: 123, agex: 23}
	// fmt.Println(p4)

	// json

	jsons, errs := json.Marshal(p1)
	fmt.Println(jsons, errs)
	fmt.Println(string(jsons))

	var res2 Person
	json.Unmarshal(jsons, &res2)

	fmt.Println(res2)
}

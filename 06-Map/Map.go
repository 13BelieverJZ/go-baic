package main

import "fmt"

func main() {
	var p1 map[int]string
	p1 = make(map[int]string)
	p1[1] = "Tom"
	fmt.Println(p1)

	var p2 map[int]string = map[int]string{}
	p2[1] = "Jack"
	fmt.Println(p2)

	var p3 map[int]string = make(map[int]string)
	p3[1] = "Dave"
	fmt.Println(p3)

	p4 := map[int]string{}
	p4[1] = "Jack"
	fmt.Println(p4)

	p5 := make(map[int]string)
	p5[1] = "Alen"
	fmt.Println(p5)

	p6 := map[int]string{
		1: "DD",
	}
	fmt.Println(p6)

}

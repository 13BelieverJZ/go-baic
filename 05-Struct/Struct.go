package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 = Person{Name: "Bob", Age: 30}
	fmt.Println(p1)

	var p2 = Person{}
	p2.Name = "Alice"
	p2.Age = 18
	fmt.Println(p2)

	p3 := Person{Name: "MJ", Age: 45}
	fmt.Println(p3)

	// 匿名结构体
	p4 := struct {
		Name string
		Age  int
	}{Name: "AI", Age: 23}
	fmt.Println(p4)
}

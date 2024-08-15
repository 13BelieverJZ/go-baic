package main

import "fmt"

func main() {
	person := []string{"Tom", "Alice", "John"}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(person), cap(person), person)

	fmt.Println("")

	for k, v := range person {
		fmt.Printf("k=%d v=%s\n", k, v)
	}

	for i := range person {
		fmt.Printf("person[%d]=%s \n", i, person[i])
	}

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]=%s \n", i, person[i])
	}

	for _, name := range person {
		fmt.Printf("%s \n", name)
	}
}

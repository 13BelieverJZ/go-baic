package main

import "fmt"

func main() {
	i := 1
	switch i {
	case 1:
		fmt.Println("i = ", 1)
	case 2:
		fmt.Println("i = ", 2)
	case 3:
		fmt.Println("i = ", 3)
	case 4, 5, 6:
		fmt.Println("i =", "4 or 5 or 6")
	default:
		fmt.Println("i =", "xxx")
	}
}

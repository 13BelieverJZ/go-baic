package main

import "fmt"

func main() {
	var age1 uint8 = 31
	var age2 = 32
	age3 := 33 // 省略var和数据类型，变量名称一定是未声明过的
	fmt.Println(age1, age2, age3)

	var age4, age5, age6 int = 31, 32, 33
	fmt.Println(age4, age5, age6)

	var name1, age7 = "James", 32
	fmt.Println(name1, age7)

	name2, age8, height := "MJ", 45, 198
	fmt.Println(name2, age8, height)
}

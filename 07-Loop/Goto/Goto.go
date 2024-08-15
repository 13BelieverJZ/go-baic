package main

import "fmt"

func main() {
	// 改变函数内代码执行顺序，不能跨函数使用

	fmt.Println("begin")

	for i := 0; i < 10; i++ {
		if i == 6 {
			goto END
		}
		fmt.Println("i=%d", i)
	}

END:
	fmt.Println("EMD")
}

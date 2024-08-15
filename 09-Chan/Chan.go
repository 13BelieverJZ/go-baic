package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")

	go func() {
		fmt.Println("go routine")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("main end")

	// 不带缓冲的通道，进和出都会阻塞。
	// 带缓冲的通道，进一次长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。

	// 声明不带缓冲的通道
	//ch1 := make(chan string)
	//
	//// 声明带10个缓冲的通道
	ch2 := make(chan string, 10)
	ch2 <- "hello"

	// 声明只读通道
	//ch3 := make(<-chan string)

	// 声明只写通道
	ch4 := make(chan<- string, 20)

	// 写入chan
	ch4 <- "a"

	// 读取chan
	val, ok := <-ch2
	fmt.Println("ch2:", val, ok)

	// 关闭chan
	//close(ch4)

	//close 以后不能再写入，写入会出现 panic
	//重复 close 会出现 panic
	//只读的 chan 不能 close
	//close 以后还可以读取数据
}

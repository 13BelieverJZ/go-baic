package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main start")
	ch := make(chan string, 2)
	go producer(ch)
	go consumer(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("main end")
}

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	ch <- "e"
	ch <- "f"
	fmt.Println("producer end")
}

func consumer(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

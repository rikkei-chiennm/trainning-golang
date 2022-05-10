package main

import (
	"fmt"
	"time"
)

/*
The select statement is used to choose from multiple send/receive channel operations.
The select statement blocks until one of the send/receive operations is ready.
If multiple operations are ready, one of them is chosen at random.
The syntax is similar to switch except that each of the case statements will be a channel operation.
*/
func main() {
	ex1 := make(chan string)
	ex2 := make(chan string)
	go example1(ex1)
	go example2(ex2)

	select {
	case s1 := <-ex1:
		fmt.Println(s1)
	case s2 := <-ex2:
		fmt.Println(s2)
	default:
		fmt.Println("error")
	}

	// muốn chạy xong cả 2 quy trình thì dùng vòng lặp for vào select x chạy từ 1 tới 2
}
func example1(ex1 chan string) {
	time.Sleep(2 * time.Second)
	ex1 <- "run ex1"
}
func example2(ex2 chan string) {
	time.Sleep(1 * time.Second)
	ex2 <- "run ex2"
}

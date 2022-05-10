package main

import (
	"fmt"
	"time"
)

// nói về timeout, thời gian chờ, sử dụng hàm timeAfter để trả về giá thì
func main() {
	taxi1 := make(chan string, 1)
	taxi2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		taxi1 <- "run taxi 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		taxi2 <- " run taxi 2 "
	}()
	select {
	case s1 := <-taxi1:
		fmt.Println(s1)
	case <-time.After(1 * time.Second):
		fmt.Println("result 1")
	}
	select {
	case s2 := <-taxi2:
		fmt.Println(s2)
	case <-time.After(3 * time.Second):
		fmt.Println("result 2")

	}
}

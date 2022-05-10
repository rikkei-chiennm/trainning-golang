package main

import (
	"fmt"
	"time"
)

// khi 2 go co tgian cho la nhu nhau thi select se lay randoom
func student(stu chan string) {
	stu <- "B19DCCN100"
}
func teacher(tea chan string) {
	tea <- "TEA0012"
}

func main() {
	tea := make(chan string)
	stu := make(chan string)
	go teacher(tea)
	go student(stu)
	time.Sleep(time.Second * 3)
	select {
	case v := <-tea:
		fmt.Println(v)
	case n := <-stu:
		fmt.Println(n)
	default:
		fmt.Println("nothing")

	}
}

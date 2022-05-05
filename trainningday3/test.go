package main

import (
	"fmt"
	"time"
)

func main() {
	go countF("sheep")
	go countF("fish")
	time.Sleep(7 * time.Second)
}
func countF(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

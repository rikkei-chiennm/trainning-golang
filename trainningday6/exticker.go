package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-ticker.C:
				fmt.Println("Ticker not stop")
			}
		}
	}()
	time.Sleep(2000 * time.Millisecond)
	done <- true
	fmt.Println("program done")
}

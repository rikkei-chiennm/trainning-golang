package main

import "fmt"

func unichan(rev chan string) {
	rev <- "hom nay buồn ngủ quá"
}
func main() {
	chal := make(chan string)
	go unichan(chal)
	fmt.Println(<-chal)
}

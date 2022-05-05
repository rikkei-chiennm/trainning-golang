package main

import "fmt"

func main() {
	var a chan int
	// khai bao channel a la nil, vì nó giống slice và map nên được khai báo theo kiểu make
	if a == nil {
		// vì là nil nên chương trình sẽ chạy vào trong if
		fmt.Println("channel a is nil, going to define it")
		// nil nên có thể khai báo theo kiểu make
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
	// Sending and receiving from a channel

	data := <-a // read from channel
	a <- data   // write to channel a

}

/*
WHAT IS CHANNEL?
- Channels can be thought of as pipes using which Goroutines communicate.
Similar to how water flows from one end to another in a pipe,
data can be sent from one end and received from the other end using channels.
DECLARING?
with channel = nil ==> name := make( chan int ) ==> declaring same to slice and maps
- Sending and receiving from a channel
Use :   "<-"
- blocking : use <- waiting return data , when return data then main running to end
==> no need time.sleep
*/

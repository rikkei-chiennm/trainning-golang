package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start program")
	ticker := time.NewTicker(400 * time.Millisecond)
	// sử dụng một ticker thời gian lặp lại là 400millis
	//cứ mỗi 400 milis lại lặp lại
	done := make(chan bool)
	// khai báo một channel tên done
	go func() {
		for {
			select {
			case check := <-done:
				if check == true {
					return
				}

			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}

	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	// dòng ticker.stop không có thì cũng không sao chương trình vẫn dừng
	done <- true
	// khi main chạy tới đây thì thằng go sẽ nhận giá trị của done là true , khi đó go sẽ return =>> kết thúc go func
	fmt.Println("Ticker stopped")
}

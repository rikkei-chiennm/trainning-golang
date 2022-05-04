package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg = sync.WaitGroup{}
	wg.Add(2)

	fmt.Println("trainning goLang day 3")

	go func() {

		count("sleep")
		wg.Done()
		//sau khi chay toi done thi chuong trinh chay xuong wg,wait

	}()

	go func() {
		count("fish")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("DONE")
	// khi có 2 gou thì chuong trình không in ra gì
	// vì bản thân main khởi tạo một gou khi đó gọi ra 2 gou kia xong thì tự đóng chương trình

	//time.Sleep(time.Second * 10)
	// khi để chương trình đợi thì bắt gou main đợi rồi chạy 2 go trên

}
func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

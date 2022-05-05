package main

import (
	"fmt"
	"time"
)

func main() {
	funcblock := make(chan bool)

	// khởi tạo channel tên funcblock trả về kiểu boolean

	fmt.Println("go main run start")
	go blocked(funcblock)
	// chạy go hàm blocked

	<-funcblock
	// chặn main bằng cách băt nó đợi cho giá trị trả về channel
	fmt.Println("main end")
}

func blocked(blockgo chan bool) {
	fmt.Println("go in func")
	time.Sleep(time.Second * 4)
	// đợi 4s
	fmt.Println("when wait done")
	blockgo <- true
	// trả dữ liệu true về kết thúc go blocked ở trong main
}

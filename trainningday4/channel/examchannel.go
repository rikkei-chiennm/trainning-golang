package main

import "fmt"

/*
Cho 1 số bất kì : 123
	bài toán : phân tích số đó ra từng chữ số : square - tổng của các bình phương từng chữ số 1*1+2*2+3*2
												cubes - tổng của các lập phương từng chữ số 1*1*1+2*2*2+3*3*3
	- XUẤT RA TỔNG CỦA CUBES VÀ QUARE
*/
func main() {
	num := 520
	// khai báo 1 số bất kì
	squarex := make(chan int)
	cubesx := make(chan int)
	// 	khai báo 2 channel
	go calcSquare(num, squarex)
	go calcCubes(num, cubesx)
	// gọi goroutine 2 hàm
	square, cubes := <-squarex, <-cubesx
	// đợi giá trị trả về từ hai hàm trên, rồi gán vào 2 biến square và cubes
	fmt.Println("Sum: ", square+cubes)
	// tính tổng 2 biến vừa được gán giá trị rồi in ra màn hình
}
func calcSquare(number int, squareOp chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		// lấy ra chữ số cuối cùng
		sum += digit * digit
		//cộng dồn dần giá trị vào sum
		number /= 10
		// loại bỏ đi chữ số cuối cùng của number
	}
	squareOp <- sum
	// viết giá trị của sum vào channel squareOp
}

//
func calcCubes(number int, cubesOp chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubesOp <- sum
	// tương tự như hàm ở trên
}

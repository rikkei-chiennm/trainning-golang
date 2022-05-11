package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// đọc file tùng dòng trong golang
func main() {

	file, err := os.Open("C:\\Users\\ChienNM\\GolandProjects\\allproject\\trainning-golang\\trainningday7\\doc.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// lúc gần kết thúc chương trình mới thực hti

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

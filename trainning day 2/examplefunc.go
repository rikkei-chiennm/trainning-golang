package main

import (
	"fmt"
	"unicode/utf8"
)

/* hàm con truyền vào 2 biến và 1 string , trả về một chuỗi con
dùng slice cắt từ ptu đầu tới ptu cuối
*/
func exam(n int, m int, o string) string {
	substring := []rune(o)
	if n > utf8.RuneCountInString(o) {
		// kiểm tra xem 2 phần tử được cắt có nằm ngoài phạm vi của mảng không
		return "out of string"
	} else {
		total := substring[n : m+1]
		return string(total)
	}

}

func main() {
	string1 := "chien-D8"
	string2 := "ｿｹｯﾄを作成する"
	n := 10
	m := 10
	fmt.Println(exam(n, m, string1))
	fmt.Println(exam(n, m, string2))
}

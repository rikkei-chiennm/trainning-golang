package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Characters are expressed by enclosing them in single quotes like this: 'a'.
	/* Trong một chuỗi string bao gồm các chữ cái không phải UTF-8
	ví dụ như: ｿｹｯﾄを作成する và một chuỗi abcdefghi.
	Số lượng ký tự của 2 chuỗi này bằng nhau nhưng Length của 2 chuỗi có khác nhau không? ==> có bằng nhau
	việc sử dụng index để lấy ra substring của 2 chuỗi này như thế nào? ==>


	*/
	myString := "Hello, 世界"
	for i, c := range myString {
		fmt.Printf("%d: %s\n", i, string(c))
	}
	fmt.Println("len of 'Hello, 世界': ", len(myString))
	fmt.Println("length of myString: ", utf8.RuneCountInString(myString))

	// Dung index de lay ra substring cua 2 chuoi ban dau

	string1 := "ｿｹｯﾄを作成する"
	stringNure := []rune(string1)

	string2 := "abcdefghi"
	for _, val := range stringNure {
		fmt.Printf("%c", val)
	}
	fmt.Println()
	for _, val1 := range string2 {
		fmt.Printf("%c", val1)
	}

}

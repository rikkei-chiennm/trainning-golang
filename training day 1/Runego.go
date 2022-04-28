package main

import "fmt"

func main() {
	/* a Rune Literal is expressed as one or more characters enclosed in single quotes like ‘g’, ‘\t’
	   In between single quotes, you are allowed to place any character except a newline and an unescaped single quote.
	*/
	var b rune
	fmt.Println(b)
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	// Displaying rune and its type
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %T", rune1,
		rune1, rune1)

	fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %T", rune2,
		rune2, rune2)

	fmt.Printf("\nRune 3: %c ;Unicode: %U; Type: %T", rune3,
		rune3, rune3)
	// creater with slice contain rune
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)
}

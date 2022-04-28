package main

import "fmt"

func main() {

	//Integers
	// integers in computer programming are whole numbers that can be positive, negative, or 0
	//default value is 0
	var n int
	fmt.Println(n)
	fmt.Println(-459)

	//Floating-Point Numbers
	//default value is 0.0f
	/*
		A floating-point number or a float is used to represent real numbers that cannot be expressed as integers.
		Real numbers include all rational and irrational numbers, and because of this, floating-point numbers can
		contain a fractional part, such as 9.0 or -116.42.
	*/

	var a float64
	fmt.Println(a)
	absoluteZero := -459.67
	fmt.Println(absoluteZero)

	// rune
	var b rune
	fmt.Println(b)
}

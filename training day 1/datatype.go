package main

import (
	"fmt"
	"unicode/utf8"
)

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

	/*
			RUNE ; default value is 0
			A rune is an alias to the int32 data type.
		    Runes represent a code point in Go.
			UTF-8 encodes all the Unicode in the range of 1 to 4 bytes, where 1 byte is used for ASCII and the rest for the Rune.
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
	/*
		BOOLEANS
		boolean data type can be one of two values, either true or false
	*/
	myBool := 5 > 8
	mybools := 5 < 7
	fmt.Println(myBool, mybools)

	/*
			STRINGS
			A string is a sequence of one or more characters (letters, numbers, symbols that can be either a constant or a variable.
		    Strings exist within either back quotes ` or double quotes "
			the backslash has no special meaning in a string literal ed: " im robot \n" => print => im robot \n
	*/
	myString := "hello to Go!"
	fmt.Println(myString)
	// create multiline strings
	newString := `This string is on 
multiple lines
within a single back 
quote on either side.`
	fmt.Println(newString)
	// Interpreted String Literals used "\ \"
	secString := "Say \"hello\" to Go!"
	fmt.Println(secString)

	// Strings with UTF-8 Characters
	a11 := "Hello, 世界z"
	for i, c := range a11 {
		fmt.Printf("%d: %s\n", i, string(c))
	}
	fmt.Println("len of 'Hello, 世界': ", len(a11))
	fmt.Println("length : ", utf8.RuneCountInString(a11))

	//Type then declaration struct

	//type student struct {
	//	name string
	//	class string

	/*
		+ Slice  ==> seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp"}
			- A slice is an ordered sequence of elements that can change in length. Slices can increase their size dynamically.
			- cap, length
			- reference
		+ Arrays ==> coral := [3]string{"blue coral", "staghorn coral", "pillar coral"}
			- An array is an ordered sequence of elements.
			- The capacity of an array is defined at creation time.
			- covetousness
		+ Maps ==> map[key]value{}
			- The map is Go’s built-in hash or dictionary type. Maps use keys and values as a pair to store data.
			- This is useful in programming to quickly look up values by an index, or in this case, a key.
			- ed: map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	*/
}

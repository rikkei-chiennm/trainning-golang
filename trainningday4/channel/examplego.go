package main

import "fmt"

func main() {

	/* vì tổng là 36 con
		tổng chân là 100 chân nên suy ra max số chó là 100/4=25
	vậy nên chó <= 25
	x là số chó, y là số gà
	*/

	y := 0
	for x := 1; x < 25; x++ {
		y = 36 - x
		sum := x*4 + y*2
		if sum == 100 {
			fmt.Printf("số chó là: %d, còn số gà là: %d", x, y)
			break
		}
	}
}

// 100 = a + b
// a bội của 2 , 4 bội của b
// x + y = 36
//chó < 25, gà < 36

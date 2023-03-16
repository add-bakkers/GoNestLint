package a

import "fmt"

func correct_simple() {
	var x int
	fmt.Scan(&x)

	if x == 0 {
		x = 1
	}
	fmt.Println("Hi ", x) // ok
}

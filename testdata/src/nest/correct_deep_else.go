package a

import "fmt"

func correct_deep_else() {
	var condition1, condition2, condition3 bool
	var x int
	if condition1 {
		x = 1
	} else {
		if condition2 { // ok
			if condition3 {
				fmt.Println("Hi!")
			}
			x = 2
		}
	}
	fmt.Println("Hi ", x)
}

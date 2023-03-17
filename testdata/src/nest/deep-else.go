package a

import "fmt"

func deep_else() {
	var condition1, condition2, condition3 bool
	var x int
	if condition1 {
		x = 1
	} else {
		if condition2 { // want `unnecessarily nested`
			if condition3 {
				fmt.Println("Hi!")
			}
		}
	}
	fmt.Println("Hi ", x)
}

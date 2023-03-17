package a

import "fmt"

func deep() {
	var condition1, condition2, condition3 bool
	var x int
	if condition1 {
		if condition2 { // want `unnecessarily nested`
			if condition3 {
				fmt.Println("Hi!")
			}
		}
		x = 1
	}
	fmt.Println("Hi ", x)
}

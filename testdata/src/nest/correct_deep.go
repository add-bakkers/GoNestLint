package a

import "fmt"

func correct_deep() {
	var condition1, condition2, condition3 bool
	var x int
	if condition1 { // ok
		if condition2 && condition3 { // ok
			fmt.Println("Hi!")
		}
		x = 1
	}
	fmt.Println("Hi ", x)
}

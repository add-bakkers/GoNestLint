package a

import "fmt"

func deep() {
	var condition1, condition2, condition3 bool
	if condition1 {
		if condition2 { // want `unnecessarily nested`
			if condition3 { // want `unnecessarily nested`
				fmt.Println("Hi!")
			}
		}

	}
}

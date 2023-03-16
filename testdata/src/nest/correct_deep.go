package a

import "fmt"

func correct_deep() {
	var condition1, condition2, condition3 bool
	if !condition1 { // ok
		return
	}
	if !condition2 { // ok
		return
	}
	if !condition3 { // ok
		return
	}
	fmt.Println("Hi!")
}

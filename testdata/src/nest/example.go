package a

import "fmt"

func simple() {
	var x int
	fmt.Scan(&x)

	if x == 0 {
		x = 1
		fmt.Println("Hi ", x)
	} else {
		fmt.Println("Hi ", x) // want `unnecessarily nested`
	}
}

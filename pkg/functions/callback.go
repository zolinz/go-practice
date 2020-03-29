package functions

import (
	"fmt"
)




func Add(x, y int) {
	fmt.Println(x + y)
}

func Multiply(x, y int) {
	fmt.Println(x * y)
}


// callback when a function is passed as an argument
func Math(x int, y int, f func(a, b int)) {
	f(x, y)
}


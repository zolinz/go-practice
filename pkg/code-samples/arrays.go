package main

import "fmt"

func main(){


	var x [5] int

	fmt.Println(x)
	x[3] =45
	fmt.Println(x)

	//slice of int
	//composite literal
	var z = [] int {1,2,3,4,5}

	fmt.Println(z)

	for i, v := range z {
		fmt.Println( i , v )
	}


	//slicing a slice

	fmt.Println(z[1:3])

}

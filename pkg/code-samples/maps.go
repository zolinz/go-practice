package main

import "fmt"

func main(){

	myMap := map [string] int {
		"Sebi":5,
		"Kati":45,
		"Zoli":45}

	 for k,v := range myMap{

	 	fmt.Println(k ,v )
	 }



	var a int = 42

	 fmt.Println( a )
	 fmt.Println(&a)

	fmt.Printf("%T\n", &a)

	b := &a // & gives you the address

	fmt.Println(b)

	fmt.Println(*b) // * gives you the value stored
}

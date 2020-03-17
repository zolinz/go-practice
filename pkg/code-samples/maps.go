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

}

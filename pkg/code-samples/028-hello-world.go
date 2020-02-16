package main

import ("fmt"
		"github.com/zolinz/go-practie/pkg/hello"
		)



func main(){

	fmt.Println("hello zolika \n", 34)

	getNames()



	fmt.Println(hello.GetZ())
}




func getNames(){
	n ,err := fmt.Println("hello from names")

	fmt.Println(n)
	fmt.Println(err)
}
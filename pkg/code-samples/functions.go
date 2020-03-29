package main

import "fmt"

func main(){


	fmt.Println("hello zolika")


	var name string

	name = "James"


	fmt.Println(sayHi(name))

}



func sayHi(msg string) string{
	fmt.Println(msg)

	lastname := "Bond"
	return lastname
}

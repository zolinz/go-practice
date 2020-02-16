package main

import (
	"fmt"
	"runtime"
)

func main(){

	var  y int = 34

	var k string = "Captain Morgan"

	fmt.Printf("%T\n Hello Zolika", y)

	fmt.Printf("%T\n", k)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)


	mystring := "hello zolika"

	bs := []byte(mystring)
	fmt.Println(bs)
	fmt.Printf("%T", bs)


}

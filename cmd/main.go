package main

import (
	"fmt"
	"github.com/zolinz/go-practie/pkg/hello"
	"os"
)




func main() {
	fmt.Println("Hello, zolika")

		fmt.Println(hello.Hello())

		fmt.Print(os.Hostname())

		 if _,err := os.Stat("/Users/zoltankovacs/zoli-int-test.config2"); os.IsNotExist(err){
			fmt.Print(err)
			panic("we are buggered")
		 }else {
			 fmt.Print(err)
		 }

		var a int = 4
		var b int32
		var c float32
		var ptr *int

		/* example of type operator */
		fmt.Printf("Line 1 - Type of variable a = %T\n", a )
		fmt.Printf("Line 2 - Type of variable b = %T\n", b )
		fmt.Printf("Line 3 - Type of variable c= %T\n", c )

		/* example of & and * operators */
		ptr = &a	/* 'ptr' now contains the address of 'a'*/
		fmt.Printf("value of a is  %d\n", a)
		fmt.Printf("*ptr is %d.\n", *ptr)



		//invoke hello using quote package
		fmt.Print(hello.Hello())




}




package main

import (
	"fmt"
	 "github.com/zolinz/go-practie/pkg/functions"
)



func main(){

	sa1 := functions.SecretAgent{
		Person: functions.Person{
			"James",
			"Bond",
		},
		Ltk: true,
	}

	sa2 := functions.SecretAgent{
		Person: functions.Person{
			"Miss",
			"Moneypenny",
		},
		Ltk: true,
	}

	p1 := functions.Person{
		First: "Zoli",
		Last:  "Kovacs",
	}

	fmt.Println(sa1)
	sa1.Speak()
	sa2.Speak()
	functions.Bar(sa1)

	functions.Bar(sa2)
	functions.Bar(p1)


	functions.MyFunction()

	//assign returning func to a variable then execute
	z := functions.MyReturningFunction()
	fmt.Println(z())

	//or
	fmt.Println(functions.MyReturningFunction()())


	//callback execution from callback.go
	functions.Math(23, 7, functions.Add)
	functions.Math(23, 7, functions.Multiply)

}





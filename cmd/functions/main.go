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

	fmt.Println(sa1)
	sa1.Speak()
	sa2.Speak()

}





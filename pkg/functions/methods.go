package functions

import "fmt"

type Person struct {
	First string
	Last string
}


type SecretAgent struct {
	Person
	Ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s SecretAgent) Speak() {
	fmt.Println("I am", s.First, s.Last)
	fmt.Println(s.Ltk)
}


func Start(){


	fmt.Print("hello zolikam")

}

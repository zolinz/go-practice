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


//interface defines method to implement
type Human interface {
	Speak()
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (p Person) Speak() {
	fmt.Println("I am", p.Last, " - the person speak")
}

// implementation is attached to a type so it becomes that interface
func (s SecretAgent) Speak() {
	fmt.Println("I am", s.First, s.Last)
	fmt.Println(s.Ltk)
}

// then that interface type can be passed in
func Bar(h Human) {
	switch h.(type) {
	case Person:
		fmt.Print("i am a Person ")
		fmt.Println("I was passed into barrrrrr", h.(Person).First)
		fmt.Println("I was passed into barrrrrr", h.(Person).First)
	case SecretAgent:
		fmt.Print("i am a Secret Agent ")
		fmt.Println("I was passed into barrrrrr", h.(SecretAgent).First)
	}
	//fmt.Println("I was passed into bar", h)
}


func Start(){


	fmt.Print("hello zolikam")

}

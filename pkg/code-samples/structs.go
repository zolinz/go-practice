package main

import "fmt"

type person struct {
	firstName string
	 lastName string
}

type secretAgent struct{
	person
	kill bool

}


func main(){
	p1 := person{"Sebi", "Kovacs"}

	fmt.Println(p1.firstName)

	var sa secretAgent

	 sa = secretAgent{p1, true}

	 fmt.Println(sa.firstName, sa.lastName )





}

package main

import (
	"encoding/json"
	"fmt"
)

type my_person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := my_person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := my_person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []my_person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

package main

import "fmt"


func getData(k chan string){
	fmt.Println("Hello from " + <- k)
}

func main()  {
	c := make(chan int)

	go func(){
		c <- 3
		fmt.Println("waiting")
	}()



	b := make(chan string)

	go getData(b)

	b <- " zolika from string channel"


	fmt.Println(<- c)
}

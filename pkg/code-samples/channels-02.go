package main

import (
	"fmt"
	"sync"
)


var wg2 sync.WaitGroup

func main()  {
	c := make(chan int)

	wg2.Add(2)

	go sendData(c)

	go recieveData(c)

	wg2.Wait()
	fmt.Println("Done ")

}


func sendData( c chan<- int){

	for i := 0; i < 100 ; i++ {
		c<-i
	}

	close(c)
	wg2.Done()
}


func recieveData( c <-chan int){

	//runtime.Gosched()
	for v := range c {
		fmt.Println(v)
	}
	//fmt.Println(<-c)
	wg2.Done()
}
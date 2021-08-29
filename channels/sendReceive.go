package main

import (
	"fmt"
	"time"
)
/*
func sendData(ch chan string)  {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func receiveData(ch chan string)  {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
	close(ch)
}


 */

func pump(ch chan int)  {
	for i := 0; ; i++{
		ch <- i
	}
}

func suck(ch chan int)  {
	for{
		fmt.Println(<-ch)
	}

}


func main()  {
	/*
	ch := make(chan string)
	go sendData(ch)
	go receiveData(ch)
	time.Sleep(1e9)
	*/

	ch1 := make(chan int)

	go pump(ch1)
	go suck(ch1)
	fmt.Println(<-ch1)
	time.Sleep(1e9)
}
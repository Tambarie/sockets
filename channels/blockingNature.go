package main

import (
	"fmt"
	"time"
)

func main()  {
	c := make(chan int)
	go func() {
		time.Sleep(5 *1e9)
		fmt.Println("received", <-c)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent",10)
}


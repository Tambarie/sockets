package main

import "fmt"

func main()  {
	// using var
	var ch1 chan string
	ch1 = make(chan string)
	fmt.Println(ch1)

	//using short operator
	ch2 := make(chan string)
	fmt.Println(ch2)

	//chanOfChans := make(chan chan string)

	funcChan := make(chan func())
	close(funcChan)
}
package main

import (
	"fmt"
	"time"
)

//type Empty interface {}
//var empty Empty
//type semaphore chan Empty

func pumpTwo() chan int{
	ch := make(chan int)
	go func() {
		for i := 0; ; i++{
			ch <- i
		}
	}()
	return ch
}

func suckTwo(ch chan int)  {
	for{
		fmt.Println(<- ch)
	}
}
func main()  {
	stream := pumpTwo()
	go suckTwo(stream)
	time.Sleep(1e9)



	//
	//var N int
	//data := make([]float64,N)
	//res := make([]float64,N)
	//sem := make(chan Empty, N) //semaphore
	//
	//for i, datum := range data {
	//	go func(i int, datum float64) {
	//		res[i] = doSomething(1,datum)
	//		sem <- empty
	//	}(i, datum)
	//}
	//for i := 0; i < N; i++ {
	//	<-sem
	//}

	//sem := make(semaphore,N)

}


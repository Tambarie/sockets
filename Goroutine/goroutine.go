package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var numCores = flag.Int("n",2,"number of CPU core to use")

func longWait(wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("End of longWait()")

}

func shortWait(wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("End of shortWait()")
}





func main()  {
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	fmt.Println("In main()")
	go longWait(wg)
	go shortWait(wg)
	wg.Wait()
	fmt.Println("About to sleep in main()")
	//time.Sleep(10 * 1e10)
	fmt.Println("At the end of main()")



}
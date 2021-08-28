package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error)  {
	if err != nil{
		fmt.Fprintf(os.Stderr,"fatal error: %s",err.Error())
		os.Exit(1)
	}
}
func main()  {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("ipv4",service)
	checkError(err)

	Listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := Listener.Accept()
		if err != nil{
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}



}
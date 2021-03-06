package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error)  {
	if err != nil{
		fmt.Fprintf(os.Stderr,"Fatal error %s\n",err.Error())
		os.Exit(1)
	}
}


func handleClient(conn net.Conn){
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil{
			return
		}

		_, err2 := conn.Write(buf[0:n])
		if err2 != nil{
			return
		}
	}
}
func main()  {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		go handleClient(conn)
	}
}
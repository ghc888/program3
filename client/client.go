package main

import (

	"fmt"
	"net"
	"program3/connectionManager"
)

const (

	CLIENT_LONG_PASSWORD uint32 = 1 << iota
	CLIENT_FILE
	CLIENT_REGISTER
)

func main() {

	con,err:=net.Dial("tcp",":9696")
	if err!=nil{
		fmt.Println("connection server error:",err)
	}

	conn,err:=connectionManager.NewConn(con)
	if err!=nil{
		fmt.Println("initial connection error:",err)
	}

	var commandFlag uint32=CLIENT_FILE
	fileName:="1.txt0"
	buffer:=make([]byte,4,4+len(fileName))


	buffer=append(buffer, byte(commandFlag))
	buffer=append(buffer,fileName...)
	fmt.Println(buffer)

	conn.Pkg.WritePacket(buffer)
	 fmt.Println("hello client")
}

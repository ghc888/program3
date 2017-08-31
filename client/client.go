package main

import (

	"fmt"
	"net"
	"program3/connectionManager"
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

	var commandFlag uint32=connectionManager.CLIENT_FILE
	fileName:="00" +
		"0"
	buffer:=make([]byte,4,4+len(fileName))

 //https://studygolang.com/articles/4350  http://blog.csdn.net/erlib
	buffer=append(buffer, byte(commandFlag))
	//十个字节的保留字段
	buffer=append(buffer,0,0,0,0,0,0,0,0,0,0)
	buffer=append(buffer,fileName...)
	fmt.Println(buffer)

	conn.Pkg.WritePacket(buffer)
	 fmt.Println("hello client")
}

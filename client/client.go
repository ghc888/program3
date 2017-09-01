package main

import (
	"fmt"
	"net"
	"program3/connectionManager"
	"program3/protocol"
	"encoding/json"
)

func RequestFileJob(conn *connectionManager.ClientConn)  {

	fileinfo := protocol.FileType{Name: "1.txt", Id:100,Path:"/tmp/"}
	message,err:=json.Marshal(fileinfo)

	if err != nil {
		fmt.Println("json marshal fileinfo error:", err)
	}

	buffer := make([]byte, 4, 1+len(message) )

	buffer = append(buffer, protocol.FILE)
	buffer = append(buffer, message...)

	conn.Pkg.WritePacket(buffer)
}
func RegisterJob(conn *connectionManager.ClientConn)  {

	register := protocol.RegisterType{Mid: "201707151127586", Project:"1500017519"}
	message,err:=json.Marshal(register)

	if err != nil {
		fmt.Println("json marshal fileinfo error:", err)
	}

	buffer := make([]byte, 4, 1+len(message) )

	buffer = append(buffer, protocol.REGISTER)
	buffer = append(buffer, message...)

	conn.Pkg.WritePacket(buffer)
}

func main() {

	con, err := net.Dial("tcp", ":9696")
	if err != nil {
		fmt.Println("connection server error:", err)
	}

	conn, err := connectionManager.NewConn(con)
	if err != nil {
		fmt.Println("initial connection error:", err)
	}


	RequestFileJob(conn)
	//RegisterJob(conn)
}

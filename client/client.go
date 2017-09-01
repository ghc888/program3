package main

import (
	"fmt"
	"net"
	"program3/connectionManager"
	"program3/protocol"
	"encoding/json"
)

func SendFileJob(conn *connectionManager.ClientConn)  {

	fileinfo := protocol.FileType{Name: "/tmp/1.txt", Size:100}
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


	//SendFileJob(conn)
	RegisterJob(conn)
}

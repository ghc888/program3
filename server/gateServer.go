package server

/*
接入控制
*/
import (


	"program3/config"
	"fmt"
	"github.com/satori/go.uuid"
	"net"
)


func GateServer(){


	config,err:=config.ParseConfigFile("/etc/1.yaml")
	if err!=nil{
		fmt.Println("parse config file error:",err)
	}

	lisAddr:=config.Addr

	tcpaddr,_:=net.ResolveTCPAddr("tcp",lisAddr)
	listener,err:=net.ListenTCP("tcp",tcpaddr)
	if err!=nil{
		fmt.Println("listen error:",err)
	}

	for{

		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("accept client request error:",err)
		}

		go HandleConn(conn)
	}
}

/*
分发client的request messge type
*/
func HandleConn(conn net.Conn) error {
	u1:=uuid.NewV4()
	fmt.Println("accept request message from client:",conn.RemoteAddr()," room id:",u1)



	return nil
}
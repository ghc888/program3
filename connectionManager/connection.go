package connectionManager

import (
	"net"
	"program3/protocol"
	"fmt"
	"sync/atomic"
	"sync"
	"log"
	"encoding/json"
)


type ClientConn struct {
	co net.Conn
	Pkg *protocol.PacketIO

	connectionId uint32

	wg         sync.WaitGroup
	mutex      sync.Mutex
}

var baseConnId uint32 = 10000

/*
初始化一个连接对象
*/
func  NewConn(co net.Conn)  (*ClientConn,error) {
	client:=new(ClientConn)

	tcpConn := co.(*net.TCPConn)
	tcpConn.SetNoDelay(false)
	client.co=tcpConn
	client.Pkg=protocol.NewPacketIO(co)

	client.Pkg.Sequence=0
	client.connectionId=atomic.AddUint32(&baseConnId, 1)
	return client,nil
}

/*
向客户端发送握手信息  15个字节:
len  message
1： 协议版本号
4： 连接id
10：保留字段
*/
func (c *ClientConn)WriteHandShake() error{

	buf:=make([]byte,4,10)

	//封装协议版本号 1个字节
	buf=append(buf, protocol.ProtocolVersion)

	//连接id 4个字节
	buf=append(buf,byte(c.connectionId),byte(c.connectionId>>8),byte(c.connectionId>>16),byte(c.connectionId>>24))

	//保留字段10个字节reserved 10 [00]
	buf=append(buf,0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	return  c.Pkg.WritePacket(buf)
}

/*
读取握手信息反馈包
4:command
*/

func (c *ClientConn)ReadHandshakeResponse()error{
	data,err:=c.Pkg.ReadPacket()
	if err!=nil{
		fmt.Println("read handshake response error:",err)
	}


	pos:=0
	command:=data[pos]
	pos++

	fmt.Println(command&protocol.REGISTER)
	fmt.Println(command&protocol.FILE)


	if command&protocol.FILE >0{
		fmt.Println("file type command")
		var file_info protocol.FileType

		err=json.Unmarshal(data[pos:],&file_info)
		if err != nil {
			log.Fatal("decode:", err)
		}
		fmt.Println(file_info)
	}else if command&protocol.REGISTER>0{
		fmt.Println("register type command")
		var register protocol.RegisterType

		err=json.Unmarshal(data[pos:],&register)
		if err != nil {
			log.Fatal("decode:", err)
		}
		fmt.Println(register)
	}


	return  nil

}


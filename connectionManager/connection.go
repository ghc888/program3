package connectionManager

import (
	"net"
	"program3/protocol"
	"fmt"
	"sync/atomic"


	"encoding/binary"
)
const (

	CLIENT_LONG_PASSWORD uint32 = 1 << iota
	CLIENT_FILE
	CLIENT_REGISTER
)


type ClientConn struct {
	co net.Conn
	Pkg *protocol.PacketIO

	connectionId uint32
	//capability uint32
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
	command:=binary.LittleEndian.Uint32(data[:4])
 	//command:=binary.LittleEndian.Uint32(data[:4])
 	if command&CLIENT_FILE>0{
		fmt.Println("file command")
	}else if command&CLIENT_REGISTER>0{
		fmt.Println("register command")
	}else {
		fmt.Println("unknown command")
	}





	pos+=4
	return  nil
}


package connectionManager

import (
	"net"
	"program3/protocol"
	"fmt"
	"encoding/binary"
	"sync/atomic"
)

type ClientConn struct {
	co net.Conn
	pkg *protocol.PacketIO

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
	client.pkg=protocol.NewPacketIO(co)

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
func (c *ClientConn)writeHandShake() error{

	buf:=make([]byte,4,10)

	//封装协议版本号 1个字节
	buf=append(buf, protocol.ProtocolVersion)

	//连接id 4个字节
	buf=append(buf,byte(c.connectionId),byte(c.connectionId>>8),byte(c.connectionId>>16),byte(c.connectionId>>24))

	//保留字段10个字节reserved 10 [00]
	buf=append(buf,0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	return  c.pkg.WritePacket(buf)
}

/*
读取握手信息反馈包
4:command
*/

func (c *ClientConn)readHandshakeResponse()error{
	data,err:=c.pkg.ReadPacket()
	if err!=nil{
		fmt.Println("read handshake response error:",err)
	}

	pos:=0
	command:=binary.LittleEndian.Uint32(data[:4])
	pos += 4
	fmt.Println("request command is:",command)
	return  nil
}

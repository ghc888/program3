package protocol

import (
	"bufio"
	"io"
	"net"
	"fmt"
	"program3/errorinfo"
)

/*
server --> client
server <-- client
交互协议
*/
const(
	MaxPacketSize  int =1<<24 - 1
	ProtocolVersion byte = 10
	TimeFormat         string = "2006-01-02 15:04:05"
	ServerVersion string = "5.6.20-program3"
)

const (
	OK_HEADER          byte = 0x00
	ERR_HEADER         byte = 0xff
	EOF_HEADER         byte = 0xfe
)

const (
	defaultReaderSize = 8 * 1024
)


type PacketIO struct {
	rb *bufio.Reader
	wb io.Writer

	Sequence uint8
}


func NewPacketIO(conn net.Conn) *PacketIO {
	p := new(PacketIO)

	p.rb = bufio.NewReaderSize(conn, defaultReaderSize)
	p.wb = conn

	p.Sequence = 0

	return p
}



func (p *PacketIO) ReadPacket() ([]byte, error) {
	header := []byte{0, 0, 0, 0}

	if _, err := io.ReadFull(p.rb, header); err != nil {
		return nil,errorinfo.ErrBadConn
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length < 1 {
		return nil, fmt.Errorf("invalid payload length %d", length)
	}

	sequence := uint8(header[3])

	if sequence != p.Sequence {
		return nil, fmt.Errorf("invalid sequence %d != %d", sequence, p.Sequence)
	}

	p.Sequence++

	data := make([]byte, length)
	if _, err := io.ReadFull(p.rb, data); err != nil {
		return nil, errorinfo.ErrBadConn
	} else {
		if length < MaxPacketSize {
			return data, nil
		}

		var buf []byte
		buf, err = p.ReadPacket()
		if err != nil {
			return nil, errorinfo.ErrBadConn
		} else {
			return append(data, buf...), nil
		}
	}
}

//data already have header
func (p *PacketIO) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPacketSize {

		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = p.Sequence

		if n, err := p.wb.Write(data[:4+MaxPacketSize]); err != nil {
			return errorinfo.ErrBadConn
		} else if n != (4 + MaxPacketSize) {
			return errorinfo.ErrBadConn
		} else {
			p.Sequence++
			length -= MaxPacketSize
			data = data[MaxPacketSize:]
		}
	}

	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.Sequence

	if n, err := p.wb.Write(data); err != nil {
		return errorinfo.ErrBadConn
	} else if n != len(data) {
		return errorinfo.ErrBadConn
	} else {
		p.Sequence++
		return nil
	}
}
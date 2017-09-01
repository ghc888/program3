package protocol

/*
server --> client
server <-- client
交互协议
0-3 :header 包头
4：type		类型
5:len 		body    body  json格式
*/
const (
	//type类型定义 长度一个字节
	FILE     = 1
	REGISTER = 2
	RQUEST   = 4
	RESPONSE = 8
	LOGOUT   = 16
	LOGIN    = 32
)

type FileType struct {
	Id int	//id 文件id标识
	Name string //文件名
	Path string //文件绝对路径名
}

type RegisterType struct {
	Mid string	 //注册的mid号
	Project string //注册的项目号
}


const (
	MaxPacketSize   int    = 1<<24 - 1
	ProtocolVersion byte   = 10
	TimeFormat      string = "2006-01-02 15:04:05"
	ServerVersion   string = "5.6.20-program3"
)

const (
	OK_HEADER  byte = 0x00
	ERR_HEADER byte = 0xff
	EOF_HEADER byte = 0xfe
)

const (
	defaultReaderSize = 8 * 1024
)

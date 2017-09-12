package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func main() {
	data:=[4]byte{0x01,0xf2,0x33,0x34}
	str:=string(data[:])
	fmt.Println(str)

	//file,err:=os.OpenFile("/tmp/1.txt",0666)
	//if err!=nil{
	//	fmt.Println("open file error:",err)
	//}

	//br:=bufio.NewReader(file)


}

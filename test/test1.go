package main

import "fmt"

func main() {
	data:=[4]byte{0x01,0xf2,0x33,0x34}
	str:=string(data[:])
	fmt.Println(str)
}

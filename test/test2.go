package main

import (
	"fmt"
	"bytes"

)

func main()  {

	ts:=[]byte("hello中国")
	r:=bytes.Runes(ts)
	fmt.Println("转换前字符串串长度:",len(ts))
	fmt.Println("转换后的字符串长度:",len(r))
	fmt.Println(ts)
	fmt.Println(r)
}
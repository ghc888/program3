package main

import (
	"fmt"
	"sync"

)
var wg sync.WaitGroup

func main()  {
	chan_n := make(chan bool)  //打印数字channel
	chan_c := make(chan bool) //打印char的channel
	wg.Add(2)

	go test(chan_n)
	go test2(chan_c)

	wg.Wait()
}

func test(status chan bool) {

	if <-status {
		for i:=0;i<10;i+=2{
			fmt.Println(i)
			fmt.Println(i+1)
			//通知字符channel
			status<-true
		}
	}

	wg.Done()
}

func test2(status chan bool)  {

	char_sequence:=[]string{"A","B","C","D","E","F","G","H","I","J"}

	if <-status{
		for i:=0;i<len(char_sequence);i+=2{
			fmt.Println(char_sequence[i])
			fmt.Println(char_sequence[i+1])
			//通知数字channel
			status<-true
		}
	}


}
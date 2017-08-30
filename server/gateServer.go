package server

import (


	"program3/config"
	"fmt"
	"github.com/satori/go.uuid"
)

func GateServer(){

	 u1:=uuid.NewV4()

	config,err:=config.ParseConfigFile("/etc/1.yaml")
	if err!=nil{
		fmt.Println("parse config file error:",err)
	}

	fmt.Println("Your listen addrr:",config.Addr," and random uuid is:",u1)
}

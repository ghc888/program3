package tools

import "github.com/satori/go.uuid"

func Getuid() uuid.UUID {
	u1:=uuid.NewV4()
	return u1
}
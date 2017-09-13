package main

import (
	"context"
	"net/http"
)

type httpServer struct {
	ctx *context.Context
	router http.Handler
}

func main() {

}
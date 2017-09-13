package main

import (
	"runtime"
	"net/http"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/",httpRouter())
	http.ListenAndServe(":5678",nil)
}

func httpRouter() (mux *http.ServeMux)  {
	mux=http.NewServeMux()
	mux.HandleFunc("/",myHandler)
	return
}

func myHandler(w http.ResponseWriter,r *http.Request)  {
	retStr:="hello http"
	retBytes:=[]byte(retStr)
	w.Write(retBytes)
}
package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"


)
type httpServer struct {
	router http.Handler
}
type APIHandler func(http.ResponseWriter, *http.Request, httprouter.Params) (interface{}, error)
type Decorator func(APIHandler) APIHandler

func Decorate(f APIHandler, ds ...Decorator) httprouter.Handle {
	decorated := f
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		decorated(w, req, ps)
	}
}
func newHTTPServer( ) *httpServer {
	router := httprouter.New()
	router.HandleMethodNotAllowed = true
	s := &httpServer{

		router: router,
	}

	//router.Handle("GET", "/ping",pingHandler)
	return s
}

func (s *httpServer) pingHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	return "OK", nil
}
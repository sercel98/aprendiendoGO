package main

import (
	"fmt"
	"net/http"
)

//Server is a struct that represents a server with a defined port
type Server struct {
	port   string
	router *Router
}

//NewServer creates a new server with an specified port
func NewServer(newPort string) *Server {
	fmt.Println("running...")
	return &Server{
		port:   newPort,
		router: NewRouter(),
	}
}

//Handle recieves a handler and assigns it to a given path and method
func (s *Server) Handle(path string, method string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler

}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

//Listen makes the server to start to listen to his port, throws error if there is any problem.
func (s *Server) Listen() error {

	http.Handle("/", s.router)

	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

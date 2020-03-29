package main

import "net/http"

//Server is a struct that represents a server with a defined port
type Server struct {
	port   string
	router *Router
}

//NewServer creates a new server with an specified port
func NewServer(newPort string) *Server {
	return &Server{
		port:   newPort,
		router: NewRouter(),
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
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

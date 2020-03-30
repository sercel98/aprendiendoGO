package main

import (
	"net/http"
)

//Router has the map of paths where a path has a map of methods
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

//NewRouter creates a new router
func NewRouter() *Router {

	//map[path]map[method]
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//FindHandler finds the corresponding handler for a path
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	//check if the path exists
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, exist, methodExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {

	handler, exist, methodExist := r.FindHandler(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)

}

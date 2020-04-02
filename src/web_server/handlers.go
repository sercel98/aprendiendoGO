package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HandleRoot is the handler for root path ("/")
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

//HandleHome is the handler for home path ("/home")
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home")
}

//HandleAPI is the handler for API path ("/api")
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API")
}

//PostRequest handles a POST request with JSON data
func PostRequest(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

//UserPostRequest handles a POST request with JSON to create an user
func UserPostRequest(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}
	response, error := user.ToJSON()

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(response)
}

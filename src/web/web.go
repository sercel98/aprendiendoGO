package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type webWritter struct {
}

//Write gest a byte slice, prints it and returns the length of the byte slice and nil(documentation)
func (webWritter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {

	response, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	e := webWritter{}
	//gets the body response and sends it to webWritter
	io.Copy(e, response.Body)
}

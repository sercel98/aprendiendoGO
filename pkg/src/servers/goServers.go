package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		channel <- server + " It's not available"
	} else {
		channel <- server + " It's running!"
	}
}

func main() {

	init := time.Now()
	channel := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://facebook.com",
		"http://instagram.com",
		"http://google.com",
	}

	for _, server := range servers {
		go checkServer(server, channel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	duration := time.Since(init)
	fmt.Printf("Execution time %s\n", duration)

}

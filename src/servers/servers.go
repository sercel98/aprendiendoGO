package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "It's not available")
	} else {
		fmt.Println(server, "It's running!")

	}
}

func main() {

	init := time.Now()

	servers := []string{
		"http://platzi.com",
		"http://facebook.com",
		"http://instagram.com",
		"http://google.com",
	}

	for _, server := range servers {
		checkServer(server)
	}
	duration := time.Since(init)
	fmt.Printf("Execution time %s\n", duration)

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//CheckLogin is the Middleware for login an user
func CheckLogin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

//CheckAuth is the Middleware used for checking if the user is authenticated
func CheckAuth() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			//Middleware logic starts here
			flag := true
			fmt.Println("Checking Auth")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

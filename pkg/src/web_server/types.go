package main

import (
	"encoding/json"
	"net/http"
)

//Middleware is a type for Handlers
type Middleware func(http.HandlerFunc) http.HandlerFunc

type Metadata interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)

}

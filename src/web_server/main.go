package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", "GET", HandleRoot)
	server.Handle("/create", "POST", PostRequest)
	server.Handle("/user", "POST", UserPostRequest)
	server.Handle("/home", "POST", server.AddMiddleware(HandleHome, CheckAuth(), CheckLogin()))
	server.Handle("/api", "GET", HandleAPI)
	server.Listen()

}

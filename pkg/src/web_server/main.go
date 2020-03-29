package main

func main() {
	server := NewServer(":5000")
	server.Handle("/", HandleRoot)
	server.Handle("/home", server.AddMiddleware(HandleHome, CheckAuth()))
	server.Handle("/api", HandleAPI)
	server.Listen()

}

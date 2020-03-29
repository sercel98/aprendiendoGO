package main

func main() {
	server := NewServer(":5000")
	server.Handle("/", HandleRoot)
	server.Handle("/home", HandleHome)
	server.Handle("/api", HandleAPI)
	server.Listen()

}

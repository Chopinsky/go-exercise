package main

func main() {
	server := CreateServer("127.0.0.1")
	server.ListenAndServe()
}

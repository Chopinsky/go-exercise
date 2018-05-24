package main

func main() {
	if err := run(); err != nil {
		println("Unable to start the server: ", err)
	}

	println("Done...")
}

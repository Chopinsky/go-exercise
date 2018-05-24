package main

func main() {
	run()
}

func run() {
	init, err := NewBlock(nil, 1)
	if err != nil {
		println("Error: ", err)
	}

	println("Hash: ", init.Hash)
}

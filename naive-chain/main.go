package main

func main() {
	init, err := NewBlock(nil, 1)
	if err != nil {
		println("Error: ", err)
	}

	println("Hash: ", init.Hash)
}

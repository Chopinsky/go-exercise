package main

// Blockchain ...
var Blockchain []Block

// ReplaceChain ...
func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

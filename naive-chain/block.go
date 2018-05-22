package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// "github.com/davecgh/go-spew/spew"
// "github.com/gorilla/mux"
// "github.com/joho/godotenv"

// Block ...
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// ValidatBlock ...
func (b *Block) ValidatBlock(childBlock *Block) bool {
	if childBlock == nil {
		return false
	}

	return true
}

// CalcHash ...
func (b *Block) CalcHash() string {
	base := string(b.Index) + b.Timestamp + string(b.BPM) + b.PrevHash
	h := sha256.New()
	h.Write([]byte(base))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) setHash() {
	b.Hash = b.CalcHash()
}

// NewBlock ...
func NewBlock(parent *Block, BPM int) (*Block, error) {
	var index int
	var prevHash string

	t := time.Now()
	if parent == nil {
		index = 0
		prevHash = ""
	} else {
		index = parent.Index + 1
		prevHash = parent.Hash
	}

	newBlock := &Block{
		Index:     index,
		Timestamp: t.String(),
		BPM:       BPM,
		Hash:      "",
		PrevHash:  prevHash,
	}

	newBlock.setHash()
	return newBlock, nil
}

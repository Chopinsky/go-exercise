package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// "github.com/joho/godotenv"

// Block ...
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// Message ...
type Message struct {
	BPM int
}

// ValidatChildBlock ...
func (b *Block) ValidatChildBlock(childBlock *Block) bool {
	if childBlock == nil {
		return false
	}

	if b.Index+1 != childBlock.Index {
		return false
	}

	if b.Hash != childBlock.PrevHash {
		return false
	}

	if childBlock.Hash != childBlock.CalcHash() {
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

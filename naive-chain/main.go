package naivechain

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

var Blockchain []Block

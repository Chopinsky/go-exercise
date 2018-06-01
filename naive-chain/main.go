package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		initBlock := &Block{
			0,
			t.String(),
			0,
			"",
			"",
		}

		spew.Dump(initBlock)
		Blockchain = append(Blockchain, *initBlock)
	}()

	if err := run(); err != nil {
		log.Fatal("Unable to start the server: ", err)
	}

	println("Done...")
}

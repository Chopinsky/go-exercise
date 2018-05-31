package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

const defaultTimeout = 10 * time.Second

func run() error {
	muxHandler := makeMuxRouter()
	addr := ":" + os.Getenv("ADDR")
	s := &http.Server{
		Addr:           addr,
		Handler:        muxHandler,
		ReadTimeout:    defaultTimeout,
		WriteTimeout:   defaultTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")

	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message
	var lastBlock Block

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	lastBlock = Blockchain[len(Blockchain)-1]
	newBlock, err := NewBlock(&lastBlock, m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	if lastBlock.ValidatChildBlock(newBlock) {
		// Create the new chain and replace the old one
		newBlockChain := append(Blockchain, *newBlock)
		ReplaceChain(newBlockChain)

		// Debugging tool
		spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	resp, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error\n"))
		return
	}

	w.WriteHeader(code)
	w.Write(resp)
}

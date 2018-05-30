package main

import (
	"net/http"
	"os"
	"time"

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

}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	serverDefaultWSPath   = "/ws"
	serverDefaultPushPath = "/push"
)

var defaultUpgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

// Server ...
type Server struct {
	Addr      string
	WSPath    string
	PushPath  string
	Upgrader  *websocket.Upgrader
	AuthToken func(token string) (userID string, ok bool)
	PushAuth  func(r *http.Request) bool
	wsHandler *websocketHandler
	ph        *pushHandler
}

// Start ..
func Start() {
	fmt.Println("All is done!")
}

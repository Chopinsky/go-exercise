package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type websocketHandler struct {
	upgrader   *websocket.Upgrader
	binder     *binder
	calcUserID func(token string) (userID string, ok bool)
}

type pushHandler struct {
	authFunc func(r *http.Request) bool
	binder   *binder
}

// RegisterMessage ...
type RegisterMessage struct {
	Token string
	Event string
}

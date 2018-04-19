package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// ErrIllegalRequest ...
var ErrIllegalRequest = errors.New("Illegal data request")

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

// PushMessage ...
type PushMessage struct {
	UserID  string `json:"userId"`
	Even    string
	Message string
}

// ServeHttp ...
func (handler *websocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	wsConn := NewConnection(conn)
	wsConn.AfterRead = func(messageType int, r io.Reader) {
		var rm RegisterMessage

		decoder := json.NewDecoder(r)
		if err := decoder.Decode(&rm); err != nil {
			return
		}

		userID := rm.Token
		if handler.calcUserID != nil {
			uid, ok := handler.calcUserID(rm.Token)
			if !ok {
				return
			}

			userID = uid
		}

		handler.binder.Bind(userID, rm.Event, wsConn)
	}

	wsConn.BeforeClose = func() {
		handler.binder.Unbind(wsConn)
	}

	wsConn.Listen()
}

func (handler *websocketHandler) closeConnections(userID, event string) (int, error) {
	connections, err := handler.binder.FilterConnections(userID, event)
	if err != nil {
		return 0, err
	}

	cnt := 0
	for i := range connections {
		// unbind
		if err := handler.binder.Unbind(connections[i]); err != nil {
			log.Printf("Unbinding failed: %v", err)
			continue
		}

		// close
		if err := connections[i].Close(); err != nil {
			log.Printf("Unable to close the connections: %v", err)
			continue
		}

		cnt++
	}

	return cnt, nil
}

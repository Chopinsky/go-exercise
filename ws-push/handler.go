package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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
	Event   string
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

// ServeHttp ...
func (s *pushHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if s.authFunc != nil {
		if ok := s.authFunc(r); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	var pm PushMessage
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&pm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrIllegalRequest.Error()))
		return
	}

	if pm.UserID == "" || pm.Event == "" || pm.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrIllegalRequest.Error()))
		return
	}

	cnt, err := s.push(pm.UserID, pm.Event, pm.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result := strings.NewReader(fmt.Sprintf("Message sent to: %d clients", cnt))
	io.Copy(w, result)
}

func (s *pushHandler) push(userID, event, message string) (int, error) {
	if userID == "" || event == "" || message == "" {
		return 0, errors.New("Input parameters can't be null")
	}

	connections, err := s.binder.FilterConnections(userID, event)
	if err != nil {
		return 0, fmt.Errorf("Failed to filter connections: %v", err)
	}

	count := 0
	for i := range connections {
		_, err := connections[i].Write([]byte(message))
		if err != nil {
			s.binder.Unbind(connections[i])
			continue
		}

		count++
	}

	return count, nil
}

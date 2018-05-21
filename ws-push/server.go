package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

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

// CreateServer ...
func CreateServer(addr string) *Server {
	return &Server{
		Addr:     addr,
		WSPath:   serverDefaultWSPath,
		PushPath: serverDefaultPushPath,
	}
}

// ListenAndServe ...
func (s *Server) ListenAndServe() error {
	b := &binder{
		userIDToConnMap:   make(map[string]*[]eventConnection),
		connIDToUserIDMap: make(map[string]string),
	}

	// request handler
	handler := websocketHandler{
		upgrader: defaultUpgrader,
		binder:   b,
	}

	if s.Upgrader != nil {
		handler.upgrader = s.Upgrader
	}

	if s.AuthToken != nil {
		handler.calcUserID = s.AuthToken
	}

	s.wsHandler = &handler
	http.Handle(s.WSPath, s.wsHandler)

	// push handler
	pushHandler := pushHandler{
		binder: b,
	}

	if s.PushAuth != nil {
		pushHandler.authFunc = s.PushAuth
	}

	s.ph = &pushHandler
	http.Handle(s.PushPath, s.ph)

	return http.ListenAndServe(s.Addr, nil)
}

// Push ...
func (s *Server) Push(userID, event, message string) (int, error) {
	return s.ph.push(userID, event, message)
}

// Drop ...
func (s *Server) Drop(userID, event string) (int, error) {
	return s.wsHandler.closeConnections(userID, event)
}

func (s *Server) check() error {
	if !checkPath(s.WSPath) {
		return fmt.Errorf("Web-Socket Path: %s is illegal", s.WSPath)
	}

	if !checkPath(s.PushPath) {
		return fmt.Errorf("Push Path: %s is illegal", s.PushPath)
	}

	if s.WSPath == s.PushPath {
		return errors.New("Web-Socket path should not be identical to the push path")
	}

	return nil
}

func checkPath(path string) bool {
	if path != "" && !strings.HasPrefix(path, "/") {
		return false
	}

	return true
}

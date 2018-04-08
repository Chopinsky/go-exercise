package main

import (
	"errors"
	"fmt"
	"sync"
)

type eventConnection struct {
	Event      string
	Connection *Connection
}

type binder struct {
	mu                sync.RWMutex
	userIDToConnMap   map[string]*[]eventConnection
	connIDToUserIDMap map[string]string
}

func (b *binder) Bind(userID string, event string, conn *Connection) error {
	if userID == "" {
		return errors.New("userID can't be blank")
	}

	if event == "" {
		return errors.New("event can't be blank")
	}

	if conn == nil {
		return errors.New("conn can't be null")
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	if eventConnections, ok := b.userIDToConnMap[userID]; ok {
		for i := range *eventConnections {
			if (*eventConnections)[i].Connection == conn {
				// already bound, we're done
				return nil
			}
		}

		newEventConnection := append(*eventConnections, eventConnection{event, conn})
		b.userIDToConnMap[userID] = &newEventConnection
	} else {
		b.userIDToConnMap[userID] = &[]eventConnection{{event, conn}}
	}

	b.connIDToUserIDMap[conn.GetID()] = userID
	return nil
}

func (b *binder) Unbind(conn *Connection) error {
	if conn == nil {
		return errors.New("Connection can't be null")
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	connID := conn.GetID()
	userID, ok := b.connIDToUserIDMap[connID]

	if !ok {
		return fmt.Errorf("can't find userID by connID: %s", connID)
	}

	if eventConns, ok := b.userIDToConnMap[userID]; ok {
		for i := range *eventConns {
			if (*eventConns)[i].Connection == conn {
				newEventConns := append((*eventConns)[:i], (*eventConns)[i+1:]...)
				delete(b.connIDToUserIDMap, connID)

				if len(newEventConns) > 0 {
					b.userIDToConnMap[userID] = &newEventConns
				} else {
					delete(b.userIDToConnMap, userID)
				}

				return nil
			}
		}

		return fmt.Errorf("Unable to locate the connection with ID: %s", connID)
	}

	return fmt.Errorf("Unable to find the event connection with user id: %s", userID)
}

func (b *binder) FindConnection(connID string) (*Connection, bool) {
	if connID == "" {
		return nil, false
	}

	userID, ok := b.connIDToUserIDMap[connID]

	if ok {
		if evtConnections, ok := b.userIDToConnMap[userID]; ok {
			conn := find(evtConnections, connID)
			if conn != nil {
				return conn, true
			}
		}

		return nil, false
	}

	for _, evtConnections := range b.userIDToConnMap {
		conn := find(evtConnections, connID)
		if conn != nil {
			return conn, true
		}
	}

	return nil, false
}

func (b *binder) FilterConnections(userID, event string) ([]*Connection, error) {
	if userID == "" {
		return nil, errors.New("User ID can't be empty")
	}

	b.mu.RLock()
	defer b.mu.RUnlock()

	if evtConnections, ok := b.userIDToConnMap[userID]; ok {
		matches := make([]*Connection, 0, len(*evtConnections))

		for i := range *evtConnections {
			if event == "" || (*evtConnections)[i].Event == event {
				matches = append(matches, (*evtConnections)[i].Connection)
			}
		}

		return matches, nil
	}

	return []*Connection{}, nil
}

func find(connArr *[]eventConnection, connID string) *Connection {
	for i := range *connArr {
		if (*connArr)[i].Connection.IsMatch(connID) {
			return (*connArr)[i].Connection
		}
	}

	return nil
}

package main

import (
	"errors"
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

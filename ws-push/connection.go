package main

import (
	"errors"
	"io"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Connection ...
type Connection struct {
	Connection   *websocket.Conn
	AfterRead    func(messageType int, reader io.Reader)
	BeforeClose  func()
	once         sync.Once
	id           string
	closeChannel chan struct{}
}

// GetID returns the id generated using UUID algorithm.
func (c *Connection) GetID() string {
	c.once.Do(func() {
		u := uuid.New()
		c.id = u.String()
	})

	return c.id
}

// IsMatch finds out if the current connection matches the connID
func (c *Connection) IsMatch(connID string) bool {
	return (c.GetID() == connID)
}

func (c *Connection) Write(frame []byte) (n int, err error) {
	select {
	case <-c.closeChannel:
		return 0, errors.New("Unable to write because the connection is closed")
	default:
		err = c.Connection.WriteMessage(websocket.TextMessage, frame)
		if err != nil {
			return 0, err
		}

		return len(frame), nil
	}
}

// Listen ...
func (c *Connection) Listen() {
	c.Connection.SetCloseHandler(func(code int, text string) error {
		if c.BeforeClose != nil {
			c.BeforeClose()
		}

		if err := c.Close(); err != nil {
			log.Println(err)
			return err
		}

		message := websocket.FormatCloseMessage(code, "")
		c.Connection.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
		return nil
	})

	if err := processConnection(c); err != nil {
		log.Println(err)
	}
}

// Close ...
func (c *Connection) Close() error {
	select {
	case <-c.closeChannel:
		return errors.New("Connection has already been closed")
	default:
		c.Connection.Close()
		close(c.closeChannel)
		return nil
	}
}

// NewConnection ...
func NewConnection(c *websocket.Conn) *Connection {
	return &Connection{
		Connection:   c,
		closeChannel: make(chan struct{}),
	}
}

func processConnection(c *Connection) error {
	for {
		select {
		case <-c.closeChannel:
			return nil
		default:
			messageType, reader, err := c.Connection.NextReader()
			if err != nil {
				return err
			}

			if c.AfterRead != nil {
				c.AfterRead(messageType, reader)
			}
		}
	}
}

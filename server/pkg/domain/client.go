package domain

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ws     *websocket.Conn
	SendCh chan []byte
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ws:     ws,
		SendCh: make(chan []byte),
	}
}

func (c *Client) Read(broadCastCh chan []byte, unregisterCh chan *Client) {
	defer func() {
		unregisterCh <- c
		c.ws.Close()
	}()

	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		broadCastCh <- msg
		log.Printf("message: %s", msg)
	}
}

func (c *Client) Write() {
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case msg, ok := <-c.SendCh:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.ws.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

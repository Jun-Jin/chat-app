package main

import (
	"chat-app-backend/pkg/domain"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	mesaageBufferSize = 256
)

// a simple REST API server listening on port 8080
// with a single endpoint /hello
// return a json response with a message "Hello, world!"
func main() {
	hub := domain.NewHub()
	go hub.Run()

	// register a handler for /hello endpoint
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ws", newWebsocketHandler(hub).Handle)

	// start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// write the response
	w.Write([]byte(`{"message": "Hello, world!"}`))
}

type websocketHandler struct {
	hub *domain.Hub
}

func newWebsocketHandler(hub *domain.Hub) *websocketHandler {
	return &websocketHandler{hub: hub}
}

func (h *websocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: socketBufferSize,
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := domain.NewClient(ws)
	go client.Read(h.hub.BroadcastCh, h.hub.UnregisterCh)
	go client.Write()
	h.hub.RegisterCh <- client
}

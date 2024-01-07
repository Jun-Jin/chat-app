package domain

type Hub struct {
	// Registered clients.
	Clients      map[*Client]bool
	RegisterCh   chan *Client
	UnregisterCh chan *Client
	BroadcastCh  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnregisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.RegisterCh:
			h.register(c)
		case c := <-h.UnregisterCh:
			h.unregister(c)
		case m := <-h.BroadcastCh:
			h.boradCast(m)
		}
	}
}

func (h *Hub) register(c *Client) {
	h.Clients[c] = true
}

func (h *Hub) unregister(c *Client) {
	if _, ok := h.Clients[c]; ok {
		delete(h.Clients, c)
		close(c.SendCh)
	}
}

func (h *Hub) boradCast(msg []byte) {
	for c := range h.Clients {
		select {
		case c.SendCh <- msg:
		default:
			h.unregister(c)
		}
	}
}

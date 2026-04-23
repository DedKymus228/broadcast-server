package client

import (
	"sync"
)

type Hub struct {
	Client     map[*Client]bool
	Broadcast  chan BroadcastMessage
	Register   chan *Client
	Unregister chan *Client
	Mu         sync.Mutex
}

type BroadcastMessage struct {
	Message  []byte
	SenderID *Client
}

func NewHub() *Hub {
	return &Hub{
		Client:     make(map[*Client]bool),
		Broadcast:  make(chan BroadcastMessage),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case user := <-h.Register:
			h.Mu.Lock()
			h.Client[user] = true
			h.Mu.Unlock()

		case user := <-h.Unregister:
			if _, ok := h.Client[user]; ok {
				h.Mu.Lock()
				delete(h.Client, user)
				h.Mu.Unlock()
				close(user.Receiver)
			}

		case msg := <-h.Broadcast:
			for c := range h.Client {
				if c == msg.SenderID {
					continue
				}

				select {
				case c.Receiver <- msg.Message:
				default:
					h.Mu.Lock()
					delete(h.Client, c)
					h.Mu.Unlock()
					close(c.Receiver)
				}
			}
		}
	}
}

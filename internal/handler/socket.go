package handler

import (
	"fmt"
	"sync"

	"golang.org/x/net/websocket"
)

type Socket struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

func NewSocket() *Socket {
	return &Socket{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (cm *Socket) AddClient(conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[conn] = true
}

func (cm *Socket) RemoveClient(conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.clients, conn)
}

func (cm *Socket) Broadcast(message string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

    fmt.Println("Broadcast called")
	for conn := range cm.clients {
		if _, err := conn.Write([]byte(message)); err != nil {
			conn.Close()
			delete(cm.clients, conn)
		}
	}
}

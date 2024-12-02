package websocket

import (
	"fmt"
	"sync"
)

type Pool struct {
	Register   chan *Client
	UnRegister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	mu         sync.Mutex
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		mu:         sync.Mutex{},
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.mu.Lock()
			pool.Clients[client] = true
			pool.mu.Unlock()
			fmt.Printf("New client joined. Total clients: %d\n", len(pool.Clients))
			pool.broadcastToAll(Message{Type: 1, Body: "New User Joined..."})
		case client := <-pool.UnRegister:
			pool.mu.Lock()
			delete(pool.Clients, client)
			pool.mu.Unlock()
			fmt.Printf("Client disconnected. Total clients: %d\n", len(pool.Clients))
			pool.broadcastToAll(Message{Type: 1, Body: "User Disconnected..."})
		case message := <-pool.Broadcast:
			fmt.Println("Broadcasting message to all clients in the pool")
			pool.broadcastToAll(message)

		}
	}
}

func (pool *Pool) broadcastToAll(message Message) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	for client := range pool.Clients {
		go func(c *Client) {
			if err := c.Conn.WriteJSON(message); err != nil {
				fmt.Printf("Error sending message to client: %v\n", err)
				pool.UnRegister <- c // Remove client with write errors
			}
		}(client)
	}
}

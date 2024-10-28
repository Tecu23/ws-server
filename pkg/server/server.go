// Package server keeps all the server logic
package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Server should handle all connections
type Server struct {
	clients      map[*websocket.Conn]bool
	clientsMutex sync.Mutex
	upgrader     websocket.Upgrader
}

// NewServer creates a server instance
func NewServer() *Server {

	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all connections for now
		},
	}

	return &Server{
		clients:      make(map[*websocket.Conn]bool),
		clientsMutex: sync.Mutex{},
		upgrader:     upgrader,
	}

}

// SetupServer creates the routes and sets server up
func (s *Server) SetupServer() {
	http.HandleFunc("/ws", s.handleConnection)
}

func (s *Server) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading connection: ", err)
		return
	}

	defer conn.Close()

	fmt.Println("New client connected!")

	// Keep listening for messages
	for {
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("Connection closed:", err)
			return
		}

		// Echo the received message back to the client
		fmt.Printf("Received: %s\n", message)
		err = conn.WriteMessage(messageType, []byte("Message received"))
	}
}

// Main is the entry point of the server
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections for now
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading connection: ", err)
		return
	}

	defer conn.Close()

	fmt.Println("New client connected!")

	// Keep listening for messages
	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("Connection closed:", err)
			break
		}

		// Echo the received message back to the client
		fmt.Printf("Received: %s\n", message)
		conn.WriteMessage(websocket.TextMessage, []byte("Message received"))
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)

	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

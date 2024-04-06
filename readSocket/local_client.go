package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Connect to WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Infinite loop to read messages from the WebSocket connection
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		// Process the received message
		message := string(msg)
		fmt.Printf("[%s] Received: %s\n", time.Now().Format(time.Stamp), message)

		// Check the message content and perform actions accordingly
		switch message {
		case "Connected":
			fmt.Printf("[%s] Status: Connected\n", time.Now().Format(time.Stamp))
			// Add your actions when status is Connected
		case "Disconnected":
			fmt.Printf("[%s] Status: Disconnected\n", time.Now().Format(time.Stamp))
			// Add your actions when status is Disconnected
		default:
			// Handle other message types if needed
		}
	}
}

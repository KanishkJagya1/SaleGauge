package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
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
			// fmt.Printf("[%s] Status: Connected\n", time.Now().Format(time.Stamp))
			if err := openChrome("https://www.google.com"); err != nil {
				log.Println("failed to open Chrome:", err)
			}
			stop := make(chan struct{})
			go printAudioStream(stop)
			// Add your actions when status is Connected
		case "Disconnected":
			fmt.Printf("[%s] Status: Disconnected\n", time.Now().Format(time.Stamp))
			// Add your actions when status is Disconnected
		default:
			// Handle other message types if needed
		}
	}
}
func printAudioStream(stop <-chan struct{}) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Generate and print random audio-like data
			timestamp := time.Now().Format(time.Stamp) // Get current timestamp in desired format
			bytesReceived := rand.Intn(151) + 100      // Generate random bytes received (100 to 250)
			latency := rand.Intn(20) + 1               // Generate random latency (1 to 20)

			// Print formatted message with timestamp, bytes received, and latency
			fmt.Printf("[%s] %d Bytes Received, Latency = %dms\n", timestamp, bytesReceived, latency)
		case <-stop:
			// Stop printing when signaled to stop
			return
		}
	}
}

func openChrome(url string) error {
	// Command to open Chrome with a specified URL
	cmd := exec.Command("google-chrome", url)
	return cmd.Run()
}

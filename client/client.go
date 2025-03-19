package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var (
	serverURL           = "ws://localhost:9090/ws"
	reconnectAttempts   = 0
	maxReconnectAttempts = 5
)

func main() {
	connectWebSocket()
}

// Handles WebSocket connection
func connectWebSocket() {
	fmt.Println("Attempting to connect to WebSocket server...")

	// Establish connection
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Println("WebSocket connection failed:", err)
		reconnect()
		return
	}
	fmt.Println(" Connected to WebSocket server!")
	// defer conn.Close()

	

	// Start goroutines for sending and receiving messages
	reconnectChan := make(chan bool)
	go readMessages(conn, reconnectChan)
	go sendMessages(conn, reconnectChan)

	<-reconnectChan
	fmt.Println(" Reconnecting...")
	reconnect()
	
}

// Handles incoming messages
func readMessages(conn *websocket.Conn ,reconnectChan chan bool) {
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket closed:", err)
			reconnectChan <- true
			return
		}
		fmt.Println("\nServer:", string(message))
	}
}

// Sends messages typed by the user
func sendMessages(conn *websocket.Conn ,reconnectChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("You: ")
		if scanner.Scan() {
			message := scanner.Text()
			if message == "exit" {
				fmt.Println("Closing connection...")
				conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Goodbye!"))
				time.Sleep(1 * time.Second)
				conn.Close()
				reconnectChan <- true
				break
			}
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Error sending message:", err)
				break
			}
		}
	}
}

// Reconnect logic
func reconnect() {
	if reconnectAttempts < maxReconnectAttempts {
		reconnectAttempts++
		fmt.Println("Reconnecting... Attempt", reconnectAttempts)
		time.Sleep(3 * time.Second)
		connectWebSocket()
	} else {
		fmt.Println("Max reconnect attempts reached. Exiting.")
	}
}

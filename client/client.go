package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {

	serverURL := url.URL{Scheme: "ws", Host: "localhost:8081", Path: "/ws"}
	log.Printf("Connecting to %s", serverURL.String())

	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Send message to server
	message := "Hello, Nisha and Ritu!!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal("Write error:", err)
	}

	// Receive echoed message
	_, reply, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Read error:", err)
	}
	log.Printf("Received: %s", reply)
}

// websocket_test.go
package main

import (
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

func startTestServer() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Failed to upgrade:", err)
			return
		}
		defer conn.Close()

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if err := conn.WriteMessage(messageType, p); err != nil {
				return
			}
		}
	})

	go func() {
		log.Println("Test WebSocket server started on :8081")
		http.ListenAndServe(":8081", nil)
	}()
}

func TestWebSocketClientServer(t *testing.T) {
	// Start the WebSocket test server
	startTestServer()

	url := url.URL{Scheme: "ws", Host: "localhost:8081", Path: "/ws"}
	t.Logf("Connecting to %s", url.String())

	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		t.Fatalf("Dial error: %v", err)
	}
	defer conn.Close()

	testMessage := "Hello, Nisha and Ritu this side!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(testMessage))
	if err != nil {
		t.Fatalf("WriteMessage error: %v", err)
	}

	// Test receiving the message back
	_, p, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("ReadMessage error: %v", err)
	}

	if string(p) != testMessage {
		t.Fatalf("Expected message: %s, but got: %s", testMessage, p)
	}

	t.Logf("Successfully received message: %s", p)
}

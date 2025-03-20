// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/gorilla/websocket"
// 	"websocket-poc/database"
// )

// func TestMain(m *testing.M) {

// 	var err error
// 	db, err = database.NewSQLiteDatabase("test_logs.db")
// 	if err != nil {
// 		log.Fatalf("Test database initialization failed: %v", err)
// 	}

// 	exitCode := m.Run()

// 	db.Close()
// 	os.Exit(exitCode)
// }
// func TestWebSocketServer(t *testing.T) {
// 	if db == nil {
// 		t.Fatal("Database is not initialized in test")
// 	}

// 	server := httptest.NewServer(http.HandlerFunc(handleConnections))
// 	defer server.Close()

// 	wsURL := "ws" + server.URL[4:] + "/ws"
// 	fmt.Println("WebSocket URL:", wsURL)

// 	ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
// 	if err != nil {
// 		t.Fatalf("WebSocket connection failed: %v", err)
// 	}
// 	defer ws.Close()

// 	testMessage := "Hello Server"
// 	err = ws.WriteMessage(websocket.TextMessage, []byte(testMessage))
// 	if err != nil {
// 		t.Fatalf("Failed to send message: %v", err)
// 	}

// 	time.Sleep(500 * time.Millisecond)

// 	_, response, err := ws.ReadMessage()
// 	if err != nil {
// 		t.Fatalf("Failed to read response: %v", err)
// 	}

// 	if string(response) != testMessage {
// 		t.Errorf("Expected response '%s', got '%s'", testMessage, response)
// 	}

// }



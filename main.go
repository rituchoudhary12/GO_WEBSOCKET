

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"github.com/gorilla/websocket"
// 	"websocket-poc/database"
// 	"time"

// )

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins
// }

// var db *database.SQLiteDatabase

// func handleConnections(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("Error upgrading connection:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	fmt.Println("Client connected")

// 	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
// 	conn.SetPongHandler(func(string) error {
// 		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
// 		return nil
// 	})

// 	for {
// 		messageType, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Client disconnected:", err)
// 			break
// 		}
// 		fmt.Printf("Received: %s\n", msg)

// 		_, err = db.StoreLog(string(msg))
// 		if err != nil {
// 			fmt.Println("Error storing log:", err)
// 		}

// 		err = conn.WriteMessage(messageType, msg)
// 		if err != nil {
// 			fmt.Println("Error writing message:", err)
// 			break
// 		}
// 	}
// }

// func main() {
// 	var err error
// 	db, err = database.NewSQLiteDatabase("logs.db")
// 	if err != nil {
// 		fmt.Println("Database initialization failed:", err)
// 		return
// 	}

// 	if db == nil {
// 		fmt.Println("Database is nil after  initialization. Exiting...")
// 		return
// 	}

// 	fmt.Println("Database initialized")
// 	defer db.Close()

// 	http.HandleFunc("/ws", handleConnections)

// 	port := ":9090"

// 	fmt.Println("WebSocket Server started on", port)
// 	err = http.ListenAndServe(port, nil)
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", message)

		if err = conn.WriteMessage(messageType, message); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", WebSocketHandler)
	log.Println("WebSocket server started at :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}


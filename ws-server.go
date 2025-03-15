package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader: Upgrades HTTP connections to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func handleWsConn(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ws upgrader failed:", err)
		return
	}
	// close connection when done
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read ws msg:", err)
			break
		}

		log.Printf("Received msg: %s\n", msg)

		// Respond back
		err = conn.WriteMessage(msgType, []byte(fmt.Sprintf(`{"event": "event_name", "data": "server:%s"}`, msg)))
		if err != nil {
			log.Println("Failed to write ws msg:", err)
			break
		}

	}
}

var (
	wsport = flag.Int("wsport", 8080, "The ws server port")
)

func WsServer() {
	http.HandleFunc("/ws", handleWsConn)
	log.Printf("WebSocket server running on ws://localhost:%d/ws\n", *wsport)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *wsport), nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}

Here's a complete guide to building a WebSocket server in Go, covering the basics and adding features similar to **Socket.IO**, including automatic reconnection, JSON handling, broadcasting, and rooms.

---

## 1. **Initializing a WebSocket Server**
Go provides WebSocket support through the `github.com/gorilla/websocket` package.

### Install Gorilla WebSocket:
```sh
go get -u github.com/gorilla/websocket
```

### Basic WebSocket Server:
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader to upgrade HTTP to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		// Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}
		fmt.Printf("Received: %s\n", msg)

		// Echo message back
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	fmt.Println("WebSocket server running on ws://localhost:8080/ws")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
```

---

## 2. **Listening to WebSocket Events**
The server can handle multiple event types by using a simple message format like JSON.

### Example: Handling JSON Events
Modify `handleConnections` to process structured JSON messages:
```go
import (
	"encoding/json"
)

// Define a struct for messages
type Message struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}

		fmt.Printf("Event: %s, Data: %s\n", msg.Event, msg.Data)

		// Echo event back
		err = conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
```

---

## 3. **Reconnecting After Client Disconnects**
Clients need to handle reconnection logic. The server itself doesn't "reconnect" but can track clients.

### Client-Side Reconnection Logic (JavaScript Example):
```javascript
let socket;

function connect() {
    socket = new WebSocket("ws://localhost:8080/ws");

    socket.onopen = () => {
        console.log("Connected to WebSocket");
    };

    socket.onmessage = (event) => {
        console.log("Received:", event.data);
    };

    socket.onclose = () => {
        console.log("Disconnected. Reconnecting in 3 seconds...");
        setTimeout(connect, 3000);
    };
}

connect();
```

---

## 4. **Handling JSON Data Efficiently**
JSON parsing and event handling help organize the WebSocket server.

### Example: Sending and Receiving JSON Messages
Modify `handleConnections` to process different events:
```go
var clients = make(map[*websocket.Conn]bool) // Track connected clients

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true // Add client to map
	defer delete(clients, conn)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Client disconnected:", err)
			break
		}

		switch msg.Event {
		case "message":
			fmt.Println("Broadcasting message:", msg.Data)
			broadcastMessage(msg)
		case "ping":
			conn.WriteJSON(Message{Event: "pong", Data: "Pong response"})
		default:
			fmt.Println("Unknown event:", msg.Event)
		}
	}
}

// Broadcast to all clients
func broadcastMessage(msg Message) {
	for client := range clients {
		client.WriteJSON(msg)
	}
}
```

---

## 5. **Broadcasting to All Clients**
This ensures that all connected clients receive updates.

### Example:
```go
func broadcastMessage(msg Message) {
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			fmt.Println("Error broadcasting:", err)
			client.Close()
			delete(clients, client)
		}
	}
}
```

---

## 6. **Implementing Rooms (Similar to Socket.IO)**
To implement rooms, use a `map[string]map[*websocket.Conn]bool`:

```go
var rooms = make(map[string]map[*websocket.Conn]bool)

func joinRoom(conn *websocket.Conn, room string) {
	if rooms[room] == nil {
		rooms[room] = make(map[*websocket.Conn]bool)
	}
	rooms[room][conn] = true
}

func leaveRoom(conn *websocket.Conn, room string) {
	delete(rooms[room], conn)
	if len(rooms[room]) == 0 {
		delete(rooms, room)
	}
}

func broadcastToRoom(room string, msg Message) {
	for client := range rooms[room] {
		client.WriteJSON(msg)
	}
}
```

Modify `handleConnections` to support joining and leaving rooms:
```go
case "join_room":
	joinRoom(conn, msg.Data)
	fmt.Println("Joined room:", msg.Data)
case "leave_room":
	leaveRoom(conn, msg.Data)
	fmt.Println("Left room:", msg.Data)
case "room_message":
	dataParts := strings.SplitN(msg.Data, ":", 2)
	if len(dataParts) == 2 {
		room := dataParts[0]
		message := dataParts[1]
		broadcastToRoom(room, Message{Event: "room_message", Data: message})
	}
```

---

## 7. **Handling Disconnections Gracefully**
Modify `handleConnections`:
```go
for {
	_, _, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Client disconnected:", err)
		delete(clients, conn)
		conn.Close()
		break
	}
}
```

---

## 8. **Handling Ping/Pong for Heartbeats**
WebSocket supports **ping/pong** to detect dead connections.

Modify the `Upgrader`:
```go
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
```

Modify `handleConnections`:
```go
conn.SetPingHandler(func(appData string) error {
	fmt.Println("Received ping")
	return conn.WriteMessage(websocket.PongMessage, []byte("pong"))
})
```

---

## 9. **Adding Middleware for Authentication**
```go
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token != "valid_token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func main() {
	http.HandleFunc("/ws", authMiddleware(handleConnections))
}
```

---

## 10. **Scaling with Multiple Servers**
Use **Redis Pub/Sub** or **NATS** for distributed WebSocket servers.

---

## **Conclusion**
This WebSocket server includes:
✅ Basic connection  
✅ Event-based messaging (like Socket.IO)  
✅ Automatic client reconnection  
✅ JSON handling  
✅ Broadcasting to all clients  
✅ Room-based messaging  
✅ Ping/pong for heartbeats  
✅ Authentication middleware  
✅ Scalability considerations  

Would you like an example of integrating this with a frontend? 🚀
package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

var server *socketio.Server
var port int

func main() {
	fmt.Println("\nHello from com.tu.ws")
}

func init() {
	initWsServer()
}

func getNestedValue(data map[string]interface{}, keys ...string) (any, bool) {
	var current interface{} = data

	for _, key := range keys {
		m, ok := current.(map[string]interface{})
		if !ok {
			return "", false
		}
		current, ok = m[key]
		if !ok {
			return "", false
		}
	}

	// Final assertion to string
	result, ok := current.(string)
	return result, ok
}

func initWsServer() {
	log.Println("Initializing ws server...")
	port = 8080
	server = socketio.NewServer(nil)

	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		log.Println("\nIO CONNECTED")
		// conn.Emit("message", "Hello from ws server")
		return nil
	})
	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		log.Println("📩 Received:", msg)
		s.Emit("message", "Server received: "+msg)
	})

	server.OnEvent("/", "json", func(s socketio.Conn, msg string) {
		log.Println("[on_json]", msg)
		var data map[string]any
		err := json.Unmarshal([]byte(msg), &data)
		if err != nil {
			log.Println("Failed to convert data to json:", err)
			return
		}
		carMake, ok := getNestedValue(data, "car", "make")
		if !ok {
			return
		}
		log.Println("make: ", carMake)
	})

	server.OnError("/", func(s socketio.Conn, err error) {
		log.Println("[on_err]", err)
	})
	server.OnDisconnect("/", func(s socketio.Conn, why string) {
		log.Println("[on_disconnect]", why)
	})
	go server.Serve()

	defer closeServer()

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)

	})
	// http.HandleFunc("/hi/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`{"hello": "world"}`))
	// })
	log.Printf("Listening on :%d/ws...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

func closeServer() {
	log.Println("\nClosing ws server...")
	server.Close()
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ServerInit() {
	fmt.Println("Initializer")
	port := findAvailablePort(45874)
	http.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		// localhost:3000/files?path=/path.to/file
		filePath := r.URL.Query().Get("path")
		fmt.Println("\nFilepath: \n", filePath)
		audioPath := filePath // Change this to your file path
		// audioPath := "C:\\Users\\User\\Music\\sample.mp3" // Windows (Use double backslashes)

		// Ensure file exists
		if _, err := os.Stat(audioPath); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		http.ServeFile(w, r, audioPath)
	})

	fmt.Printf("Serving audio at http://localhost:%d/files\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
func ScndServer() {
	// Create a route handler
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {

		structResp := struct {
			Hello string `json:"hello"`
		}{Hello: "World"}
		jsonResp, _ := json.Marshal(structResp)

		w.Write(jsonResp)
	})

	http.ListenAndServe(":5555", nil)
}

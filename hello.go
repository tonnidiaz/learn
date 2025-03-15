package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func vars() {
	var age int = 24
	fmt.Println("His age is:", age)

	var netWorth float32 = 500000.65
	fmt.Println("Net worth:", netWorth)

	auto := "This is an auto-typed variable" // Can only be used inside funcs
	fmt.Println("Auto:", auto)

	var check_type = "John Doe"
	fmt.Printf("check_type has a value of: %v, and a type of: %T\n", check_type, check_type)
}
func hello_main() {
	_init()
	fmt.Println("Hello Tonni Diaz")
	vars()

}

// Check if a port is available and return the first open port
func findAvailablePort(startPort int) int {
	for port := startPort; port <= startPort+10; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			ln.Close()
			return port
		}
	}
	fmt.Println("No available ports found in range.")
	os.Exit(1)
	return 0
}

func _init() {
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

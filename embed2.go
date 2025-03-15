package main

import (
	"embed"
	"fmt"
)

//go:embed all:bin2
var binFS embed.FS

func embed2_main() {
	fmt.Println("\nEmbed2 main")

	fmt.Println("Embed2 list files")
	entries, err := binFS.ReadDir("bin2")
	if err != nil {
		fmt.Println("\nEmbed2 error:", err)
		return
	}

	for i, entry := range entries {
		fmt.Printf("[%d] %s\n", i, entry.Name())
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("\nHello from [tu-go]")
	args := os.Args
	// fmt.Println("Running with args: [", args[1:], "]")
	if len(args) >= 2 {
		name := args[1]
		fmt.Printf("Hello to you [%s]\n", name)
	} else {
		fmt.Println("No name specified")
	}

}

package main

import "flag"

var (
	gport = flag.Int("port", 50051, "The server port")
)

func GRPCServer() {
	println("\nThe port is:", *gport)
}

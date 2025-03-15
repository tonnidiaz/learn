package main

import "time"

func longProcess() {
	println("Starting long process...")
	time.Sleep(5 * time.Second)
	println("Long process done!")
}

func afterLongProcess() {
	println("\nAfter long process")
}
func asyncFunc() {
	go longProcess()
	afterLongProcess()
}

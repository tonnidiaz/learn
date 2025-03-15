package main

import (
	"encoding/json"
	"fmt"
)

// Example function that takes a *[]string
func getFileSz(files *[]string) {
	fmt.Println("Files:", *files)
}

func conv_main() {
	jsonStr := `["file1.mp3", "file2.mp3", "file3.mp3"]`

	var data any // or interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Type assertion to convert data to []interface{}
	fileInterface, ok := data.([]any)
	if !ok {
		fmt.Println("Error: data is not an array")
		return
	}

	// Convert []interface{} to []string
	// files := make([]string, len(fileInterface))
	// for i, v := range fileInterface {
	// 	files[i], ok = v.(string)
	// 	if !ok {
	// 		fmt.Println("Error: file is not a string")
	// 		return
	// 	}
	// }
	files := ToSlice[string](&fileInterface)
	// Pass a pointer to the []string
	getFileSz(&files)
}

func ToSlice[T any](inp *[]any) []T {
	_interface := *inp
	out := make([]T, len(_interface))
	for i, v := range _interface {
		out[i] = v.(T)
	}
	return out
}

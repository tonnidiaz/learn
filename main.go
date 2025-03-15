package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"runtime"
	"slices"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	tonni := Person{Name: "Tonni", Age: 24}
	fmt.Println("Main.go")
	fmt.Println(tonni)
	// embed_main()
	// embed2_main()
	// asyncFunc()
	// ServerInit()
	// StartMPV("localhost")
	// AnotherServer()
	WsServer()

}

func init() {
	// ExecCmd()
}

func ExecCmd() {
	argsStr := `ls` //`keytool -genkeypair -v -keystore /media/tonni/win/src/Mint/Documents/RF/Mobile/Keystores/tu-keystore.keystore -alias tu-key0 -keyalg RSA -keysize 2048 -validity 10000 -dname "CN=Tonni Diaz, OU=Tu-unit, O=Tu-org, L=JHB, ST=Gauteng, C=SA" -storepass StorePass444 -keypass KeyPass444`
	var sh string
	if runtime.GOOS == "windows" {
		sh = "cmd /C"
	} else {
		sh = "sh -c"
	}

	fmt.Println("Args:", argsStr)
	cmd := exec.Command(strings.Split(sh, " ")[0], strings.Split(sh, " ")[1], argsStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Printf("{pid: %d}\n", cmd.Process.Pid)
	fmt.Println("Output:\n", string(output))

	jsonString := `{"name":"Tonni Diaz","age":20}`
	var j map[string]any
	json.Unmarshal([]byte(jsonString), &j)
	fmt.Println("JSOD_DATA:", j["name"])
}
func SomeFunc() {
	exts := []string{"mp3", "ogg", "mp4"}
	filename := "file.ts"
	splitFilename := strings.Split(filename, ".")
	slices.Reverse(splitFilename)
	ext := splitFilename[0]
	fmt.Println(exts, ext)
	fmt.Printf("Exts contain %s? %v\n", ext, slices.Contains(exts, ext))
}

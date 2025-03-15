package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
)

var videoPath2 string = "/media/tonni/win/Windows.old/Users/squas/AppData/Local/Startup/Include/Config/Config/miami_swim.mkv"

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func startMPV() {
	cmd := exec.Command("mpv", "--no-terminal", "--no-audio", "--vo=tct", "--frames=1000", videoPath2)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Failed to start MPV:", err)
	}
}

func streamRGBFrames(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Failed:", err)
		return
	}
	defer conn.Close()

	// Start MPV and send RGB frames over WebSocket
	cmd := exec.Command("ffmpeg", "-i", videoPath2, "-f", "rawvideo", "-pix_fmt", "rgb24", "pipe:1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error getting stdout:", err)
		return
	}
	err = cmd.Start()
	if err != nil {
		log.Println("Failed to start FFmpeg:", err)
		return
	}

	buf := make([]byte, 800*450*3) // RGB24 frame size
	for {
		_, err := stdout.Read(buf)
		if err != nil {
			break
		}
		conn.WriteMessage(websocket.BinaryMessage, buf)
	}
	cmd.Wait()
}

func main() {
	http.HandleFunc("/stream", streamRGBFrames)
	go startMPV()

	log.Println("WebSocket server running on ws://localhost:8080/stream")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"

	mpv "github.com/blang/mpv"
)

var filePath string = "/media/tonni/win/Windows.old/Users/squas/AppData/Local/Startup/Include/Config/Config/miami_swim.mkv"
var filePath2 string = "/home/tonni/Downloads/King of the Hill S01-S13 (1997-)/King of the Hill S01 (360p re-dvdrip)/King of the Hill S01E01 Pilot.mp4"

func StartMPV(ip string) {

	println("\nStarting MPV...")

	ipcc := mpv.NewIPCClient("/tmp/mpvsocket") // Lowlevel client
	player := mpv.NewClient(ipcc)

	streamURL := "http://127.0.0.1:8080" // or your desired IP/port
	err := player.SetProperty("stream-lavf-o", streamURL)
	if err != nil {
		fmt.Println("Error setting stream output:", err)
	}

	// Load file with streaming
	err = player.Loadfile(filePath2, mpv.LoadFileModeReplace)
	if err != nil {
		fmt.Println("Error starting stream:", err)
	}
}

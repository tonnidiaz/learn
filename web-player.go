package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"

	"github.com/gorilla/mux"
)

var (
	mpvCmd    *exec.Cmd
	ffmpegCmd *exec.Cmd
	mu        sync.Mutex
)

func AnotherServer() {

	fmt.Println("\nAnotherServer...")
	r := mux.NewRouter()

	r.HandleFunc("/stream", streamHandler).Queries("seek", "{seek}")
	r.HandleFunc("/stream", streamHandler)
	r.PathPrefix("/live/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.StripPrefix("/live/", http.FileServer(http.Dir("./live"))).ServeHTTP(w, r)
	}))

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	mu.Lock()
	defer mu.Unlock()

	seekTime := r.URL.Query().Get("seek")
	println("\nSeeking to:", seekTime)
	stopStream()

	startStream(filePath, seekTime)

	http.Redirect(w, r, "/live/stream.m3u8", http.StatusFound)
}

func startStream(videoPath, seekTime string) {
	if err := createDirIfNotExist("./live"); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	mpvPath, err := exec.LookPath("mpv")
	if err != nil {
		log.Fatalf("mpv not found: %v", err)
	}

	ffmpegPath, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatalf("ffmpeg not found: %v", err)
	}

	mpvArgs := []string{
		"--no-cache",
		"--no-terminal",
		"--quiet",
		"--vo=null",
		"--ao=null",
		videoPath,
		"-o", "-",
	}
	if seekTime != "" {
		mpvArgs = append(mpvArgs, "--start="+seekTime)
	}

	mpvCmd = exec.Command(mpvPath, mpvArgs...)

	ffmpegCmd = exec.Command(ffmpegPath,
		"-i", "-",
		"-c:v", "libx264",
		"-preset", "veryfast",
		"-tune", "zerolatency",
		"-c:a", "aac",
		"-ar", "44100",
		"-f", "hls",
		"-hls_time", "2",
		"-hls_list_size", "0",
		"-hls_segment_filename", "./live/stream%03d.ts",
		"./live/stream.m3u8",
	)
	ffmpegCmd.Stderr = log.Writer() // requires import "log"

	ffmpegCmd.Stdin, _ = mpvCmd.StdoutPipe()
	err = mpvCmd.Start()
	if err != nil {
		log.Printf("Error starting mpv: %v", err)
		return
	}
	err = ffmpegCmd.Start()
	if err != nil {
		log.Printf("Error starting ffmpeg: %v", err)
		return
	}

	log.Println("MPV and FFmpeg stream started.")
}

func stopStream() {
	if mpvCmd != nil && mpvCmd.Process != nil {
		mpvCmd.Process.Kill()
		mpvCmd.Wait()
		log.Println("MPV stopped.")
	}
	if ffmpegCmd != nil && ffmpegCmd.Process != nil {
		ffmpegCmd.Process.Kill()
		ffmpegCmd.Wait()
		log.Println("FFmpeg stopped.")
	}
}

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

package main

// import (
// 	"embed"
// 	"fmt"
// 	"io"
// 	"io/fs"
// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"runtime"
// )

// // Test embeding files/binaries

// //go:embed all:bin
// var binFs embed.FS

// const TEMP_DIRNAME = "tu-player-temp-45874587"

// var BIN_DIR = filepath.Join(os.TempDir(), TEMP_DIRNAME, "bin")
// var tempFfmpegFilePtn = "ffmpeg-*"

// /*
//  - Check if ffmpeg bin exists in temp_dirname
//  - if !exist:
// 	- open and copy the bundled ffmpeg bytes to temp_dirname/ffmpeg
// 	- Close the bundled ffmpeg file when done
//  - exec commands
// */

// func embed_main() {

// 	defer fmt.Print("\nAll done dawg!\n")

// 	file, err := binFs.Open("bin/ffmpeg")
// 	if err != nil {
// 		fmt.Println("Unable to Bin dir")
// 		return
// 	}
// 	defer closeFile(&file)

// 	stat, err := file.Stat()
// 	if err != nil {
// 		fmt.Print("Error reading file stat", err)
// 		return
// 	}
// 	fmt.Println("\nTu bin content")
// 	fmt.Println("Name:", stat.Name())
// 	fmt.Println("Size:", stat.Size())
// 	execFfmpegCmd()
// }

// func execBin() {
// 	cmd := exec.Command("bin/tu-bin")
// 	// cmd.Run()
// 	output, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("Could not get output", err)
// 		return
// 	}
// 	fmt.Println("\nOutput:", string(output))
// }

// func closeFile(file *fs.File) {
// 	fmt.Printf("Closing file...")
// 	(*file).Close()
// }

// func execFfmpegCmd() {

// 	fmt.Println("\nCreating ffmpeg file...")
// 	execFilename, err := createFfmpegBinary()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Printf("\nExecuting ffmpeg stuff on [%s]\n", execFilename)

// 	cmd := exec.Command(execFilename, "-version")
// 	// cmd.Run()
// 	out, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("\nError while executing FFMPEG", err)
// 		return
// 	}

// 	fmt.Println("\nCommand output:\n", string(out))
// }

// var appName = "tu-player-45874587"

// func getFFmpegPath() (string, error) {
// 	var appDir string

// 	switch runtime.GOOS {
// 	case "windows":
// 		appDir = filepath.Join(os.Getenv("LOCALAPPDATA"), appName, "bin")
// 	case "darwin": // macOS
// 		appDir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", appName, "bin")
// 	default: // Linux and others
// 		appDir = filepath.Join(os.Getenv("HOME"), fmt.Sprintf(".%s", appName), "bin")
// 	}

// 	// Ensure the directory exists
// 	if err := os.MkdirAll(appDir, 0755); err != nil {
// 		return "", err
// 	}

// 	// Return the full path to ffmpeg binary
// 	ffmpegPath := filepath.Join(appDir, "ffmpeg")
// 	if runtime.GOOS == "windows" {
// 		ffmpegPath += ".exe"
// 	}
// 	return ffmpegPath, nil
// }

// func createFfmpegBinary() (string, error) {
// 	ffmpegFile, err := binFs.Open("bin/ffmpeg")
// 	if err != nil {
// 		return "", err
// 	}

// 	// Close file when done
// 	defer ffmpegFile.Close()

// 	outPath, err := getFFmpegPath()
// 	if err != nil {
// 		return "", err
// 	}

// 	ffmpegExecFile, err := os.Open(outPath)
// 	if err != nil {
// 		// Create temp file
// 		fmt.Println("Creating output file due to", err)
// 		ffmpegExecFile, err = os.Create(outPath)
// 		if err != nil {
// 			return "", err
// 		}
// 		// Copy contents of bundled bin to temp bin
// 		fmt.Println("Copying binaries...")
// 		_, err = io.Copy(ffmpegExecFile, ffmpegFile)
// 		if err != nil {
// 			return "", err
// 		}

// 	}
// 	defer (*ffmpegExecFile).Close()

// 	// Make the temp file executable
// 	fmt.Println("Making temp binaries executable...")
// 	if err := os.Chmod(ffmpegExecFile.Name(), 0755); err != nil {
// 		return "", err
// 	}

// 	fmt.Println("Done creating binaries!")
// 	return ffmpegExecFile.Name(), nil
// }

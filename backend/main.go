package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"terje-configurator/pkg/api"
	"terje-configurator/pkg/db"
)

//go:embed dist/*
var content embed.FS

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}

func main() {
	// Default to current directory + TerjeSettings if not provided
	defaultPath := "./TerjeSettings"
	if abs, err := filepath.Abs(defaultPath); err == nil {
		defaultPath = abs
	}

	// If running from inside backend folder during dev, adjust
	if _, err := os.Stat(defaultPath); os.IsNotExist(err) {
		// Try to look up one level (useful for dev)
		defaultPath = "../TerjeSettings"
	}

	rootPath := flag.String("path", defaultPath, "Path to TerjeSettings directory")
	port := flag.String("port", "8080", "Port to run server on")
	flag.Parse()

	absPath, _ := filepath.Abs(*rootPath)
	fmt.Printf("Starting Terje Configurator...\n")
	fmt.Printf("Root Path: %s\n", absPath)

	// Verify if root path exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Printf("WARNING: Root path does not exist: %s\n", absPath)
		fmt.Printf("Please ensure the TerjeSettings folder is present.\n")
	}

	api.SetRootPath(absPath)

	// Init DB
	dbPath := filepath.Join(".", "terje_history.db")
	if err := db.InitDB(dbPath); err != nil {
		log.Fatalf("Failed to init DB: %v", err)
	}
	fmt.Println("Database initialized.")

	// Handlers
	http.HandleFunc("/api/tree", api.HandleTree)
	http.HandleFunc("/api/file", api.HandleFile)
	http.HandleFunc("/api/history", api.HandleHistory)
	http.HandleFunc("/api/restore", api.HandleRestore)

	// Serve Embedded Frontend
	distFS, err := fs.Sub(content, "dist")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(distFS)))

	fmt.Printf("Server running on http://localhost:%s\n", *port)

	// Open browser in a goroutine to not block startup
	go func() {
		openBrowser("http://localhost:" + *port)
	}()

	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatal(err)
	}
}

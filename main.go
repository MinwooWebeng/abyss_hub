package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const staticDir = "./static"

func fileHandler(w http.ResponseWriter, r *http.Request) {
	// Sanitize and construct the full file path
	fmt.Println(r.URL.Path)
	requestedPath := filepath.Clean(r.URL.Path)
	fullPath := filepath.Join(staticDir, requestedPath)

	// Check if it's a file (not a directory) and it exists
	info, err := os.Stat(fullPath)
	if err != nil || info.IsDir() {
		http.NotFound(w, r)
		return
	}

	// Serve the file
	http.ServeFile(w, r, fullPath)
}

func main() {
	http.HandleFunc("/", fileHandler)

	log.Println("Serving static files on http://localhost:7070")
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

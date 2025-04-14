package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
)

//go:embed static/*
var staticFiles embed.FS

// serveStaticFile serves static files, including index.html.
func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	// If the request is for '/', serve the index.html file
	if r.URL.Path == "/" {
		data, err := staticFiles.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "Index not found", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
		return
	}

	// Otherwise, try to serve the requested static file
	data, err := staticFiles.ReadFile("static" + r.URL.Path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Set the appropriate content-type based on the file extension
	switch path.Ext(r.URL.Path) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".html":
		w.Header().Set("Content-Type", "text/html")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	w.Write(data)
}

func main() {
	// Serve static files from the embedded "static" folder
	http.HandleFunc("/", serveStaticFile)

	// Start the server
	fmt.Println("Server started on :8080 🚀")
	http.ListenAndServe(":8080", nil)
}

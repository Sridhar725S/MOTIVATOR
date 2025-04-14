package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

//go:embed static/*
var embeddedFiles embed.FS

var quotes = []string{
	"You got this! 💪",
	"Keep going, code slayer!",
	"One bug at a time.",
	"Commit like it's hot 🔥",
	"Break problems, not keyboards 💻✊",
	"Drink coffee, write code ☕💻",
	"Push it real good 🚀",
	"Sleep is for the weak – just kidding, go rest 😴",
	"Your future self will thank you 👨‍💻✨",
	"One PR away from greatness 🧠",
	"Code like nobody's debugging 👀",
	"Every line brings you closer to done 🏁",
	"Refactor like a rockstar 🎸",
	"Trust the process – and Git 🌀",
	"Keep calm and console.log() 📟",
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS and Content-Type headers
	w.Header().Set("Access-Control-Allow-Origin", "https://motivator.onrender.com")
	w.Header().Set("Content-Type", "application/json")
	random := quotes[rand.Intn(len(quotes))]
	// Respond with a random quote
	json.NewEncoder(w).Encode(map[string]string{"quote": random})
}

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	// Get the file path from the URL
	filePath := r.URL.Path
	// Read the corresponding file from embedded files
	fileData, err := embeddedFiles.ReadFile("static" + filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Set correct MIME type based on file extension
	switch {
	case strings.HasSuffix(filePath, ".css"):
		w.Header().Set("Content-Type", "text/css")
	case strings.HasSuffix(filePath, ".js"):
		w.Header().Set("Content-Type", "application/javascript")
	case strings.HasSuffix(filePath, ".html"):
		w.Header().Set("Content-Type", "text/html")
	case strings.HasSuffix(filePath, ".json"):
		w.Header().Set("Content-Type", "application/json")
	default:
		w.Header().Set("Content-Type", "application/octet-stream") // default for unknown file types
	}

	// Serve the static file content
	w.Write(fileData)
}

func main() {
	// Serve static files from embedded folder under "/static/"
	http.HandleFunc("/static/", serveStaticFile)

	// Serve quotes API at "/api/quote"
	http.HandleFunc("/api/quote", quoteHandler)

	// Serve the index.html file at the root ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := embeddedFiles.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "Index not found 💀", http.StatusInternalServerError)
			return
		}
		// Set the Content-Type for HTML and write the index file
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	// Start the server on port 8080
	fmt.Println("Server is started on :8080 🚀")
	http.ListenAndServe(":8080", nil)
}

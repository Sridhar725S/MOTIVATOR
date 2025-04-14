package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

//go:embed static/**/*
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
	w.Header().Set("Access-Control-Allow-Origin", "https://motivator.onrender.com")
	w.Header().Set("Content-Type", "application/json")
	random := quotes[rand.Intn(len(quotes))]
	json.NewEncoder(w).Encode(map[string]string{"quote": random})
}

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	// Get the file path
	filePath := r.URL.Path[len("/static/"):] // Remove "/static/" from the URL path
	fileData, err := embeddedFiles.ReadFile("static/" + filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Set correct content-type for file
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
		w.Header().Set("Content-Type", "application/octet-stream") // default
	}

	// Serve the file content
	w.Write(fileData)
}

func main() {
	// Serve static files from the embedded folder
	http.HandleFunc("/static/", serveStaticFile)

	// Serve quotes API
	http.HandleFunc("/api/quote", quoteHandler)

	// Serve index.html at root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := embeddedFiles.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "Index not found 💀", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	// Start the server
	fmt.Println("Server is started on :8080 🚀")
	http.ListenAndServe(":8080", nil)
}

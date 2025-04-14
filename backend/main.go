package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"math/rand"
	"net/http"
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
	w.Header().Set("Access-Control-Allow-Origin", "https://motivator-73sm.onrender.com")
	w.Header().Set("Content-Type", "application/json")
	random := quotes[rand.Intn(len(quotes))]
	json.NewEncoder(w).Encode(map[string]string{"quote": random})
}

func main() {
	// Serve static files (like main.js, CSS, etc.)
	staticFS, _ := fs.Sub(embeddedFiles, "static")
	fsHandler := http.FileServer(http.FS(staticFS))

	// Handle API
	http.HandleFunc("/api/quote", quoteHandler)

	// Serve frontend files and fallback to index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "static/index.html")
			return
		}
		fsHandler.ServeHTTP(w, r)
	})

	fmt.Println("🚀 Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

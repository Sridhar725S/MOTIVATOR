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
	w.Header().Set("Access-Control-Allow-Origin", "https://motivator-73sm.onrender.com")
	w.Header().Set("Content-Type", "application/json")
	random := quotes[rand.Intn(len(quotes))]
	json.NewEncoder(w).Encode(map[string]string{"quote": random})
}


func main() {
	// Strip the "static/" prefix so paths match like /static/js/... in requests
	staticFS, _ := fs.Sub(embeddedFiles, "static")
	fsHandler := http.FileServer(http.FS(staticFS))

	// Quotes API
	http.HandleFunc("/api/quote", quoteHandler)

	// Handle static files like JS, CSS, JSON, etc.
	http.Handle("/static/", http.StripPrefix("/static/", fsHandler))

	// Fallback for SPA: Serve index.html for other routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || !strings.HasPrefix(r.URL.Path, "/api/") {
			data, err := embeddedFiles.ReadFile("static/index.html")
			if err != nil {
				http.Error(w, "💥 Index not found", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html")
			w.Write(data)
			return
		}
		http.NotFound(w, r) // return 404 for other missing routes
	})

	fmt.Println("Server is started on :8080 🚀")
	http.ListenAndServe(":8080", nil)
}

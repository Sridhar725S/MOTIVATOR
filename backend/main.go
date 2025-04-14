package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS 

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
	// Serve static files under /static/
	staticFS := http.FileServer(http.FS(staticFiles))
	http.Handle("/static/", http.StripPrefix("/static/", staticFS))

	// Serve index.html for root or unknown routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := "static/index.html"
		data, err := staticFiles.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Index file not found 💥", 500)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write(data)
	})

	http.HandleFunc("/api/quote", quoteHandler)

	fmt.Println("🚀 Motivator 3000 server running on :8080")
	http.ListenAndServe(":8080", nil)
}

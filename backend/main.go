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
	// Serve the static files, and redirect non-file requests to index.html
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(staticFiles))))
	http.HandleFunc("/api/quote", quoteHandler)
	fmt.Println("Server is started")
	http.ListenAndServe(":8080", nil)
}

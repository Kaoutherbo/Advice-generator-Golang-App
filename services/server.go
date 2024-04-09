package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "math/rand"
    "time"
)

type Quote struct {
    ID      int    `json:"id"`
    Content string `json:"content"`
}

// Global variable to hold all quotes
var quotes []Quote

// Function to read quotes from JSON file
func readQuotesFromFile(filename string) ([]Quote, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    err = json.NewDecoder(file).Decode(&quotes)
    if err != nil {
        return nil, err
    }

    return quotes, nil
}

// Function to get a random quote
func getRandomQuote() Quote {
    rand.Seed(time.Now().UnixNano())
    return quotes[rand.Intn(len(quotes))]
}

// Handler to serve all quotes
func getAllQuotes(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(quotes)
}

// Handler to serve a random quote
func getQuote(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    quote := getRandomQuote()
    json.NewEncoder(w).Encode(quote)
}

// enable CORS
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
        w.Header().Set("Access-Control-Allow-Methods", "https://kaoutherbo.github.io/Advice-generator-Golang-App/")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            return
        }

        next.ServeHTTP(w, r)
    })
}

func main() {
    // Read quotes from JSON file
    var err error
    quotes, err = readQuotesFromFile("data.json")
    if err != nil {
        log.Fatal("Error reading quotes:", err)
    }

    http.Handle("/", http.FileServer(http.Dir(".")))

    http.HandleFunc("/api/allQuotes", getAllQuotes)
    http.HandleFunc("/api/quote", getQuote)

    // Add CORS middleware
    http.ListenAndServe(":8080", enableCORS(http.DefaultServeMux))

    // Start the server
    fmt.Println("Server started on port 8080")
}

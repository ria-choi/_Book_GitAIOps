package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World from Notiflex!")
}

func main() {
    // Triggering CI/CD
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

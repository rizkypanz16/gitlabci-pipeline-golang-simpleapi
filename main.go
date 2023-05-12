package main

import (
    "fmt"
    "encoding/json"
    "log"
    "net/http"
)


// Initialization
type Message struct {
    Status string `json:"status"`
    Code string `json:"code"`
    Messages string `json:"message"`
    Api string `json:"api"`
}

func main() {
    fmt.Println("server listening at http://localhost:8080")
    http.HandleFunc("/", dataHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
    // JSON DATA
    message := Message{
        Status: "true",
        Code: "200",
        Messages: "SUCCESS",
        Api: "== GOLANG SIMPLE API ==",
    }

    // Set Content-Type: application/json
    w.Header().Set("Content-Type", "application/json")

    // JSON OUTPUT
    json.NewEncoder(w).Encode(message)
}
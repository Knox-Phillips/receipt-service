package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "strings"
)

type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price            string `json:"price"`
}

type Receipt struct {
    Retailer      string    `json:"retailer"`
    PurchaseDate  string    `json:"purchaseDate"`
    PurchaseTime  string    `json:"purchaseTime"`
    Items         []Item    `json:"items"`
    Total         string    `json:"total"`
}

type ReceiptResponse struct {
    ID string `json:"id"`
}

var receipts = make(map[string]Receipt)
var receiptPoints = make(map[string]int)

func processReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt Receipt

    err := json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    receiptID := generateID()

    points := calculatePoints(receipt)

    receipts[receiptID] = receipt
    receiptPoints[receiptID] = points

    response := ReceiptResponse{ID: receiptID}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
    receiptID := strings.TrimPrefix(r.URL.Path, "/receipts/")
    points, exists := receiptPoints[receiptID]
    if !exists {
        http.Error(w, "Receipt ID not found", http.StatusNotFound)
        return
    }

    response := map[string]int{"points": points}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func generateID() string {
    return fmt.Sprintf("%x", rand.Int63())
}

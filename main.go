package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Receipt represents a single receipt
type Receipt struct {
	ID           string `json:"id,omitempty"`
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
	Points       int    `json:"points,omitempty"`
}

// Item represents a single item in the receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// In-memory storage for receipts
var receipts = make(map[string]Receipt)

func main() {
	http.HandleFunc("/receipts/process", processReceipt)
	http.HandleFunc("/receipts/", getPoints)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler for POST /receipts/process
func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	receipt.ID = id
	receipt.Points = calculatePoints(receipt)
	receipts[id] = receipt

	fmt.Printf("Receipt processed and stored with ID: %s\n", id) // Debugging line
	fmt.Printf("Current stored receipts: %v\n", receipts)        // Debugging line

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// Handler for GET /receipts/{id}/points
func getPoints(w http.ResponseWriter, r *http.Request) {
	// Split the path and get only the ID part
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	id := parts[2] // Extracts just the ID portion

	// Find the receipt by ID
	receipt, exists := receipts[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		fmt.Printf("Receipt with ID %s not found\n", id) // Debugging line
		return
	}

	fmt.Printf("Found receipt with ID %s: %v\n", id, receipt) // Debugging line

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": receipt.Points})
}

// calculatePoints calculates the points for a given receipt based on the rules
func calculatePoints(receipt Receipt) int {
	points := 0

	// 1 point per alphanumeric character in the retailer name
	points += countAlphanumeric(receipt.Retailer)

	// 50 points if the total is a round dollar amount
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if int(total*100)%25 == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// 6 points if the purchase date is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the purchase time is between 2:00 pm and 4:00 pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		points += 10
	}

	// Points based on item description length
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2)) // rounding up
		}
	}

	return points
}

// countAlphanumeric returns the number of alphanumeric characters in a string
func countAlphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if isAlphanumeric(r) {
			count++
		}
	}
	return count
}

// isAlphanumeric checks if a rune is alphanumeric
func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Item represents an item with name, price, and ID.
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items []Item

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Define API endpoints
	r.HandleFunc("/items", GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", GetItem).Methods("GET")
	r.HandleFunc("/items", CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

// GetItems returns a list of all items.
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetItem returns an item by ID.
func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// CreateItem creates a new item.
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// UpdateItem updates an existing item by ID.
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedItem Item
	_ = json.NewDecoder(r.Body).Decode(&updatedItem)

	for i, item := range items {
		if item.ID == params["id"] {
			items[i] = updatedItem
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// DeleteItem deletes an item by ID.
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range items {
		if item.ID == params["id"] {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

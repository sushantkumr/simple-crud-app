package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItems(t *testing.T) {
	// Create a request to retrieve all items
	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request to the handler
	GetItems(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the response body and check the result
	var items []Item
	if err := json.NewDecoder(rr.Body).Decode(&items); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	// Validate items as needed.
}

func TestGetItem(t *testing.T) {
	// Create a request to retrieve a specific item by ID
	itemID := "1" // Replace with the desired item ID
	req, err := http.NewRequest("GET", "/items/"+itemID, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request to the handler
	GetItem(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the response body and check the result
	var item Item
	if err := json.NewDecoder(rr.Body).Decode(&item); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	// Validate item as needed.
}

func TestCreateItem(t *testing.T) {
	// Create a new item to be added
	newItem := Item{
		ID:    "4", // Replace with a unique ID
		Name:  "New Item",
		Price: 9.99,
	}

	// Marshal the item to JSON
	newItemJSON, err := json.Marshal(newItem)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request to create a new item
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(newItemJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request to the handler
	CreateItem(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the response body and check the result
	var createdItem Item
	if err := json.NewDecoder(rr.Body).Decode(&createdItem); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	// Validate createdItem as needed.
}

func TestUpdateItem(t *testing.T) {
	// Create an updated item
	updatedItem := Item{
		ID:    "1", // Replace with the ID of the item you want to update
		Name:  "Updated Item",
		Price: 19.99,
	}

	// Marshal the item to JSON
	updatedItemJSON, err := json.Marshal(updatedItem)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request to update the item
	req, err := http.NewRequest("PUT", "/items/"+updatedItem.ID, bytes.NewBuffer(updatedItemJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request to the handler
	UpdateItem(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the response body and check the result
	var updatedItemResponse Item
	if err := json.NewDecoder(rr.Body).Decode(&updatedItemResponse); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
	// Validate updatedItemResponse as needed.
}

func TestDeleteItem(t *testing.T) {
	// Create a request to delete an item by ID
	itemID := "2" // Replace with the ID of the item you want to delete
	req, err := http.NewRequest("DELETE", "/items/"+itemID, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Serve the request to the handler
	DeleteItem(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

func TestMain(m *testing.M) {
	// Before running tests, you can set up your database or server here if needed.
	// You can also tear down resources after running tests.
	// For simplicity, we didn't include database setup/teardown in this example.
	m.Run()
}

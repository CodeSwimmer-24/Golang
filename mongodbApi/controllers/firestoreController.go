package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CodeSwimmer-24/mongodbApi/config" // Use the actual module name
)

// Define the structure of the data you expect to receive
type Data struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

// HandlePostData handles the POST request and writes data to Firestore
func HandlePostData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Parse the incoming JSON data
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Reference a Firestore collection
	collection := config.FirestoreClient.Collection("your_collection_name")

	// Add a new document with the parsed data
	_, _, err := collection.Add(ctx, map[string]interface{}{
		"field1": data.Field1,
		"field2": data.Field2,
	})
	if err != nil {
		log.Printf("Error adding document: %v", err)
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	// Send a response back to the client
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Data successfully posted to Firestore")
}

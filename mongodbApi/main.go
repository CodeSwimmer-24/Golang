package main

import (
	"log"
	"net/http"

	"github.com/CodeSwimmer-24/mongodbApi/config"
	"github.com/CodeSwimmer-24/mongodbApi/routers"
)

func main() {
	// Initialize Firebase
	config.InitFirebase()
	defer config.FirestoreClient.Close()

	// Register routes
	routers.RegisterRoutes()

	// Start the server
	log.Println("Server starting at :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

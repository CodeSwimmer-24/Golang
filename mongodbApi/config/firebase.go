package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func InitFirebase() {
	// Initialize the context
	ctx := context.Background()

	// Load Firebase service account key
	sa := option.WithCredentialsFile("config/serviceAccountKey.json")

	// Set Firebase project ID
	conf := &firebase.Config{ProjectID: "gherdekho-a9354"} // Replace with your actual project ID

	// Initialize the Firebase app
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v\n", err)
	}

	// Initialize Firestore client
	FirestoreClient, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v\n", err)
	}
}

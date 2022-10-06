package database

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func GetFirebaseClient() (*firestore.Client, error) {

	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("GOOGLE_FIREBASE_CREDENTIALS_PATH"))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing firestore client: %v", err)
	}

	return client, nil
}

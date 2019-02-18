package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/option"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func getUser(ctx context.Context, app *firebase.App) *auth.UserRecord {
	uid := "D2TtllL0rM"

	// [START get_user_golang]
	// Get an auth client from the firebase.App
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %#v\n", u)
	// [END get_user_golang]
	return u
}

func main() {
	jsonPath := os.Getenv("FIREBASE_KEYFILE_PATH")
	opt := option.WithCredentialsFile(jsonPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	getUser(context.Background(), app)
}

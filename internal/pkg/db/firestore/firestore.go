package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
	"os"
)

func ConnectToFirebase() ( *firestore.Client, context.Context, error) {
	_, err := os.Stat("sa.json")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	sa := option.WithCredentialsFile("sa.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}


	return client, ctx, nil

}
package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func InitializeFirebase() *db.Client {
	ctx := context.Background()

	conf := &firebase.Config{
		DatabaseURL: "https://fyfirman-tech-default-rtdb.asia-southeast1.firebasedatabase.app/",
	}

	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	return client
}

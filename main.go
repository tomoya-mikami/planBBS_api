package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "planbbs")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	server, err := InitializeServer(client, ctx)
	if err != nil {
		log.Fatalln(err)
	}

	server.Start()
}

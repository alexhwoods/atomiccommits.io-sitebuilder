package db

import (
	"context"
	"log"

	"cloud.google.com/go/bigtable"
)

var Client *bigtable.Client

func InitDB(project string, instance string) {
	client, err := bigtable.NewClient(context.Background(), project, instance)

	if err != nil {
		log.Fatalf("Could not create data operations client: %v", err)
	}

	Client = client
}

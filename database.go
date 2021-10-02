package main

import (
	"context"
	"github.com/solher/arangolite/v2"
	"log"
)

type Node struct {
	arangolite.Document
}

var db *arangolite.Database

func handleDatabase() {

	ctx := context.Background()

	// We declare the database definition.
	db = arangolite.NewDatabase(
		arangolite.OptEndpoint("http://localhost:8529"),
		arangolite.OptBasicAuth("root", "arango"),
		arangolite.OptDatabaseName("_system"),
	)

	// The Connect method does two things:
	// - Initializes the connection if needed (JWT authentication).
	// - Checks the database connectivity.
	if err := db.Connect(ctx); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"log"

	"app/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=postgres port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	log.Printf("success")
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

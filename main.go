package main

import (
	"context"
	"log"
	"net/http"

	"example.com/enumeg/entgen"
	"example.com/enumeg/entgen/ogent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create ent client.
	client, err := entgen.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	// Run the migrations.
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}
	// Start listening.
	srv, err2 := ogent.NewServer(ogent.NewOgentHandler(client))
	if err2 != nil {
		log.Fatal(err2)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}

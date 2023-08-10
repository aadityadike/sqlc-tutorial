package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"reflect"

	"github.com/aadityadike/sqlc-tutorial/tutorial"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func run() error {

	godotenv.Load()

	DBURL := os.Getenv("DBURL")

	if DBURL == "" {
		log.Fatal("DATABASE URL DOESN'T EXIST")
	}

	ctx := context.Background()

	db, err := sql.Open("postgres", DBURL)
	if err != nil {
		return err
	}


	queries := tutorial.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err, "error")
	}
}

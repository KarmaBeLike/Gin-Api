package repo

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/pkg/errors"
)

func Init(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id uuid PRIMARY KEY,
        username text UNIQUE NOT NULL,
        email text UNIQUE NOT NULL,
        password_hash bytea NOT NULL
    );`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	preparedQuery, err := db.PrepareContext(ctx, query)
	if err != nil {
		return errors.Wrap(err, "problem in sql syntax")
	}

	_, err = preparedQuery.ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "problem execute query")
	}

	// Document
	queryD := `
	CREATE TABLE IF NOT EXISTS documents (
		id SERIAL PRIMARY KEY,
		title TEXT UNIQUE,
		content TEXT, 
		expiry_date timestamp
	);`

	_, err = db.ExecContext(ctx, queryD)
	if err != nil {
		// нужно вместо log.Printf писать return errors.Wrap(err, "Error creating documents table")
		log.Printf("Error creating documents table: %v", err)
		return err
	}

	log.Println("Documents table created or already exists")

	return nil
}

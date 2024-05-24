package repo

import (
	"context"
	"database/sql"
	"log"
	"time"
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
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	// Document
	queryD := `
	CREATE TABLE IF NOT EXISTS documents (
		id UUID PRIMARY KEY,
		title TEXT UNIQUE,
		content TEXT
	);`

	_, err = db.ExecContext(ctx, queryD)
	if err != nil {
		log.Printf("Error creating documents table: %v", err)
		return err
	}
	log.Println("Documents table created or already exists")

	// Проверяем структуру таблицы documents
	checkDocumentsTableStructure(db)

	return nil
}

func checkDocumentsTableStructure(db *sql.DB) {
	rows, err := db.Query("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = 'documents'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	log.Println("Columns in documents table:")
	for rows.Next() {
		var columnName, dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			log.Fatal(err)
		}
		log.Printf("Column Name: %s, Data Type: %s\n", columnName, dataType)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

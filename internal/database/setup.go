package database

import (
	"database/sql"
	"log"
)

// CheckAndCreateTable проверяет существование таблицы и создает ее, если она не существует.
func CheckAndCreateTable(db *sql.DB) error {
	// SQL запрос для создания таблицы, если она не существует
	createQuery := `
    CREATE TABLE IF NOT EXISTS documents (
        id UUID PRIMARY KEY,
        title TEXT UNIQUE,
        content TEXT
    );`

	_, err := db.Exec(createQuery)
	if err != nil {
		log.Printf("Error creating documents table: %v", err)
		return err
	}

	log.Println("Table 'documents' checked/created successfully.")
	return nil
}

package document

import (
	"context"
	"database/sql"
	"fmt"

	"Gin-Api/internal/model"
)

type DocumentRepository struct {
	db *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{
		db: db,
	}
}

func (dr *DocumentRepository) CreateDoc(ctx context.Context, document *model.Document) error {
	fmt.Println("document", document, "________________________")
	query := `INSERT INTO documents (title, content, expiry_date)            
              VALUES ($1, $2, $3)`
	_, err := dr.db.ExecContext(ctx, query, document.Title, document.Content, document.ExpiryDate)
	if err != nil {
		return err
	}
	return nil
}

func (dr *DocumentRepository) GetDoc(ctx context.Context, title string) (*model.Document, error) {
	query := `
	SELECT id, title, content
	FROM documents
	WHERE title = $1;`
	document := &model.Document{}
	err := dr.db.QueryRowContext(ctx, query, title).Scan(
		&document.ID,
		&document.Title,
		&document.Content,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrDocumentNotFound
		}
		return nil, err
	}
	return document, nil
}

package document

import (
	"context"
	"database/sql"

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
	query := `INSERT INTO documents (id, title, content,expiry_date)            
              VALUES ($1, $2, $3, $4)
              RETURNING id, title, content` // вставляет новый документ в таблицу. возвращаем для проверки,что данные были вставлены корректно, и эти данные могут быть использованы в дальнейшем в приложении без необходимости выполнять дополнительный запрос для их получения.
	_, err := dr.db.ExecContext(ctx, query, document.ID, document.Title, document.Content)
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
	err := dr.db.QueryRowContext(ctx, query, title).Scan( // выполняет запрос, подставляя значения document.ID, document.Title и document.Content в параметры $1, $2 и $3
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

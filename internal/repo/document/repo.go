package document

import (
	"context"
	"database/sql"

	"Gin-Api/internal/dto"
	"Gin-Api/internal/model"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type DocumentRepository struct {
	db *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{
		db: db,
	}
}

func (dr *DocumentRepository) CreateDocument(ctx context.Context, request *dto.CreateDocumentRequest) (*model.Document, error) {
	query := `INSERT INTO documents (id, title, content)            
              VALUES ($1, $2, $3)
              RETURNING id, title, content` // вставляет новый документ в таблицу. возвращаем для проверки,что данные были вставлены корректно, и эти данные могут быть использованы в дальнейшем в приложении без необходимости выполнять дополнительный запрос для их получения.

	document := &model.Document{ // создает экземпляр документа
		ID:      uuid.New().String(),
		Title:   request.Title,
		Content: request.Content,
	}

	err := dr.db.QueryRowContext(ctx, query, document.ID, document.Title, document.Content).Scan(&document.ID, &document.Title, &document.Content) // выполняет запрос, подставляя значения document.ID, document.Title и document.Content в параметры $1, $2 и $3
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "document_title_key"` { // .Scan(&document.ID, &document.Title, &document.Content) считывает возвращенные значения id, title и content и сохраняет их в соответствующие поля структуры document.
			return nil, model.ErrDuplicateTitle
		}
		return nil, err
	}

	return document, nil
}

func (dr *DocumentRepository) GetDocument(ctx context.Context, title string) (*model.Document, error) {
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

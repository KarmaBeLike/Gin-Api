package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"Gin-Api/internal/dto"
	"Gin-Api/internal/model"
)

type DocumentService struct {
	documentStorage DocumentStorage
}
type DocumentStorage interface {
	CreateDoc(context.Context, *model.Document) error
	GetDoc(context.Context, string) (*model.Document, error)
}

func NewDocumentService(docStorage DocumentStorage) *DocumentService {
	return &DocumentService{
		documentStorage: docStorage,
	}
}

func (ds *DocumentService) CreateDocument(ctx context.Context, request *dto.CreateDocumentRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	document := &model.Document{
		Title:      request.Title,
		Content:    request.Content,
		UserID:     request.UserID,
		ExpiryDate: time.Now().Add(7 * 24 * time.Hour),
	}

	err := ds.documentStorage.CreateDoc(ctx, document)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DocumentService) GetDocument(req *dto.GetDocumentRequest) (*model.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	document, err := ds.documentStorage.GetDoc(ctx, req.Title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrDocumentNotFound
		}
		return nil, err
	}
	return document, nil
}

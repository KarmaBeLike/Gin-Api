package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"Gin-Api/internal/dto"
	"Gin-Api/internal/model"
	"Gin-Api/internal/repo/document"
)

type DocumentService struct {
	docRepo *document.DocumentRepository
}

func NewDocumentService(documentRepository *document.DocumentRepository) *DocumentService {
	return &DocumentService{
		docRepo: documentRepository,
	}
}

func (ds *DocumentService) CreateDocument(ctx context.Context, request *dto.CreateDocumentRequest) (*model.Document, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	createdDoc, err := ds.docRepo.CreateDocument(ctx, request)
	if err != nil {
		return nil, err
	}

	return createdDoc, nil
}

func (ds *DocumentService) GetDocument(req *dto.GetDocumentRequest) (*model.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	document, err := ds.docRepo.GetDocument(ctx, req.Title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrDocumentNotFound
		}
		return nil, err
	}
	return document, nil
}

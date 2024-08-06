package dto

import uuid "github.com/google/uuid"

type CreateDocumentRequest struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"user_id"`
}

type GetDocumentRequest struct {
	Title string `json:"title"`
}

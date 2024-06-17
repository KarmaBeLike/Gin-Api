package dto

import uuid "github.com/google/uuid"

type CreateDocumentRequest struct { // для запроса на создание документа
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"user_id"`
}

type GetDocumentRequest struct {
	Title string `json:"title"`
}

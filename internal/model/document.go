package model

import (
	"errors"
	"time"

	uuid "github.com/google/uuid"
)

var (
	ErrDuplicateTitle   = errors.New("document with this title already exists")
	ErrDocumentNotFound = errors.New("document not found")
)

type Document struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserID     uuid.UUID `json:"user_id"`
	ExpiryDate time.Time `json:"expiry_date"`
}

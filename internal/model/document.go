package model

import "errors"

var (
	ErrDuplicateTitle   = errors.New("document with this title already exists")
	ErrDocumentNotFound = errors.New("document not found")
)

type Document struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

package model

import (
	"errors"

	uuid "github.com/google/uuid"
)

var (
	ErrDuplicateEmail     = errors.New("duplicate email")
	ErrPasswordNotCorrect = errors.New("password not correct")
	ErrRecordNotFound     = errors.New("record not found")
)

type User struct { // закидываем в бд, внутренняя структура
	ID             uuid.UUID `json:"id"`
	UserName       string    `json:"username" binding:"required" validate:"min=8,containsany=!@#?*"`
	Email          string    `json:"email" binding:"required,email"`
	HashedPassword []byte    `json:"password" binding:"required" validate:"min=8"`
	Password       string    `json:"-"`
}

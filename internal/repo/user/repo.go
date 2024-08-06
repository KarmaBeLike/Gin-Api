package user

import (
	"context"
	"database/sql"

	"Gin-Api/internal/dto"
	"Gin-Api/internal/model"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, request *dto.RegistrationRequest, hashPassword []byte) error {
	query := `INSERT INTO users(id, username, email, password_hash) VALUES ($1,$2,$3,$4)`

	user := model.User{
		ID:             uuid.New(),
		UserName:       request.UserName,
		Email:          request.Email,
		HashedPassword: hashPassword,
	}

	_, err := ur.db.ExecContext(ctx, query, user.ID, user.UserName, user.Email, user.HashedPassword)
	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return errs.Wrap(pqErr, "user is exist")
			}

		// 	if pqErr.Code == ""
		// }

		// if errors.As(err, sql) {
		// }
	}
	return nil
}

func (ur *userRepository) GetUser(ctx context.Context, username string) (*model.User, error) {
	query := `SELECT id, username, email, password_hash FROM users WHERE username=$1;`
	user := &model.User{}

	err := ur.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.HashedPassword,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrUserNotFound // Определите эту ошибку в вашем пакете model
		}
		return nil, err
	}

	return user, nil
}
// func switchError(err error) error {
// 	switch {
// 	case errors.Is(err, sql.ErrNoRows):

// 	case errors.Is(err, sql.ErrTxDone):

// 	}
// }

package user

import (
	"context"
	"time"

	"Gin-Api/internal/dto"
	"Gin-Api/internal/model"
	"Gin-Api/pkg"

	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	CreateUser(context.Context, *dto.RegistrationRequest, []byte) error
	GetUser(context.Context, string) (*model.User, error)
}
type UserService struct {
	userStorage UserStorage
}

func NewUserService(userStorage UserStorage) *UserService { // получили возможность создать user и отдаем возможность хэшировать пароль и отдаем возможность создать user
	return &UserService{
		userStorage: userStorage,
	}
}

func (us *UserService) RegisterUser(request *dto.RegistrationRequest) error {
	hash, err := pkg.HashPassword(request.Password)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = us.userStorage.CreateUser(ctx, request, hash)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) LoginUser(req *dto.LoginRequest) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	user, err := us.userStorage.GetUser(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(*req.Password))
	if err != nil {
		return nil, model.ErrPasswordNotCorrect
	}
	user.Password = *req.Password
	return user, nil
}

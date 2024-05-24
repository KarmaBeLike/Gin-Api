package user

import (
	"errors"
	"net/http"

	"Gin-Api/config"
	"Gin-Api/internal/dto"
	"Gin-Api/internal/model"

	"github.com/gin-gonic/gin"
)

type userService interface {
	RegisterUser(*dto.RegistrationRequest) error
	LoginUser(*dto.LoginRequest) (*model.User, error)
}
type UserClient struct {
	service userService
}

func NewUserClient(service userService) *UserClient {
	return &UserClient{
		service: service,
	}
}

func (c *UserClient) Routes(r *gin.Engine, cfg *config.Config) {
	r.POST("signup", c.SignUp)
	r.POST("signin", c.SignIn)
}

func (c *UserClient) SignUp(ctx *gin.Context) {
	var request dto.RegistrationRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.service.RegisterUser(&request)
	if err != nil {
		if errors.Is(err, model.ErrDuplicateEmail) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "a user with this email address already exists"})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User succcessfully registered"})
}

func (c *UserClient) SignIn(ctx *gin.Context) {
	var request dto.LoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = c.service.LoginUser(&request)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound), errors.Is(err, model.ErrPasswordNotCorrect):
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email or password"})
			return
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User successful logged in"})
}

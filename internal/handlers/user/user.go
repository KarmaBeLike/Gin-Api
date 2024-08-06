package user

import (
	"errors"
	"net/http"
	"time"

	"Gin-Api/config"
	"Gin-Api/internal/dto"
	"Gin-Api/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type userService interface {
	RegisterUser(*dto.RegistrationRequest) error
	LoginUser(*dto.LoginRequest) (*model.User, error)
	CreateToken(*model.User) string
}
type UserClient struct {
	service userService
	redis   *redis.Client
}

func NewUserClient(service userService, redisClient *redis.Client) *UserClient {
	return &UserClient{
		service: service,
		redis:   redisClient,
	}
}

func (c *UserClient) Routes(r *gin.Engine, cfg *config.Config) {
	r.POST("signup", c.SignUp)
	r.POST("signin", c.SignIn)

	r.POST("/test", c.BasicAuthMiddleware(), func(c *gin.Context) {
		// Если middleware пропускает запрос, отправляем ответ с сообщением об успешном доступе
		c.JSON(http.StatusOK, gin.H{"message": "you have access"})
	})
}

func (c *UserClient) SignUp(ctx *gin.Context) {
	var request dto.RegistrationRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	err = c.service.RegisterUser(&request)
	if err != nil {
		if errors.Is(err, model.ErrDuplicateEmail) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "a user with this email address already exists"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
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

	user, err := c.service.LoginUser(&request)
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
	token := c.service.CreateToken(user)
	c.redis.Set(token, user.ID, time.Minute)
	ctx.JSON(http.StatusOK, gin.H{"message": "User successful logged in"})
}

func (c *UserClient) TestHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Test auth"})
}

package handlers

import (
	"FIXit/backend/internal/apierr"
	"FIXit/backend/internal/config"
	"FIXit/backend/internal/http/middleware"
	"FIXit/backend/internal/repository"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

var userStore *repository.UserRepository

func getUserStore(cfg *config.Config) *repository.UserRepository {
	if userStore == nil {
		userStore = repository.NewUserRepository(cfg)
		return userStore
	}
	return userStore
}

type RegisterRequest struct {
	Name       string `json:"name" binding:"required,min=2,max=50,alpha"`
	Surname    string `json:"surname" binding:"required,min=2,max=50,alpha"`
	Patronymic string `json:"patronymic" binding:"omitempty,min=2,max=50,alpha"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8,max=64"`
	Password2  string `json:"password2" binding:"required,eqfield=Password"`
}

func Register(c *gin.Context) {
	cfg := c.MustGet("config").(*config.Config)
	var req RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		middleware.AbortErr(c, apierr.Wrap(apierr.ErrInvalidJSON, err))
		return
	}
	context := context.Background()

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err = getUserStore(cfg).Create(context, repository.User{
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
		Email:      req.Email,
		Password:   string(encryptedPassword),
		RoleID:     1,
		IsActive:   true,
	})
	if err != nil {
		middleware.AbortErr(c, apierr.Wrap(apierr.ErrInternal, err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Создан юзер " + req.Name,
	})
}

func GetUserById(c *gin.Context) {
	cfg := c.MustGet("config").(*config.Config)
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		middleware.AbortErr(c, apierr.Wrap(apierr.ErrBadRequest, err))
		return
	}
	ctx := context.Background()

	user, _ := getUserStore(cfg).FindById(ctx, int(id))
	c.JSON(http.StatusOK, user)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func Login(c *gin.Context) {
	cfg := c.MustGet("config").(*config.Config)
	var req LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		middleware.AbortErr(c, apierr.Wrap(apierr.ErrInvalidJSON, err))
		return
	}
	ctx := context.Background()
	user, err := getUserStore(cfg).FindByEmail(ctx, req.Email)
	if err != nil {
		middleware.AbortErr(c, apierr.Wrap(apierr.ErrUserNotFound, err))
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		middleware.AbortErr(c, apierr.Wrap(apierr.ErrInvalidPassword, err))
		return
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": user.ID,
		"iat": now.Unix(),
		"exp": now.Add(cfg.JWTTTL).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(cfg.JWTSecret))

	if err != nil {
		middleware.AbortErr(c, apierr.Wrap(apierr.ErrTokenCreation, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": signed,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

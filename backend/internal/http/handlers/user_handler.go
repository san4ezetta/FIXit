package handlers

import (
	"FIXit/backend/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userStore *repository.UserStore

func getUserStore() *repository.UserStore {
	if userStore == nil {
		userStore = repository.NewUserStore()
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
	var req RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = getUserStore().Create(repository.User{
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
		Email:      req.Email,
		Password:   req.Password,
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Создан юзер " + req.Name,
	})
}

func User(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, _ := getUserStore().FindById(int(id))
	c.JSON(http.StatusOK, user)
}

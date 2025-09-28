package main

import (
	"FIXit/backend/internal/http/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST("/register", handlers.Register)
	server.GET("/user/:id", handlers.User)
	err := server.Run(":8080")
	if err != nil {
		return
	}
}

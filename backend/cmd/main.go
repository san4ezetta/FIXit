package main

import (
	"FIXit/backend/internal/config"
	"FIXit/backend/internal/http/handlers"
	"FIXit/backend/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

func WithConfig(cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("config", &cfg)
		ctx.Next()
	}
}

func main() {
	cfg := config.Load()
	server := gin.Default()
	server.Use(WithConfig(cfg))
	server.POST("/register", handlers.Register)
	server.GET("/user/:id", middleware.NewAuth(cfg.JWTSecret).RequireAuth(), handlers.GetUserById)
	server.POST("/login", handlers.Login)
	err := server.Run(cfg.Addr)
	if err != nil {
		return
	}
}

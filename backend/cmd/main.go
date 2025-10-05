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
	server.Use(middleware.Responder())
	api := server.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.GET("/user/:id", middleware.NewAuth(cfg.JWTSecret).RequireAuth(), handlers.GetUserById)
		api.POST("/login", handlers.Login)
	}
	err := server.Run(cfg.Addr)
	if err != nil {
		return
	}
}

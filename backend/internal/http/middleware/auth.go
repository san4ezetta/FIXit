package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	secret []byte
}

func NewAuth(secret string) *Auth {
	return &Auth{secret: []byte(secret)}
}

func (m *Auth) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(auth, "Bearer ")

		claims := jwt.MapClaims{}
		parsed, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return m.secret, nil
		})

		if err != nil || !parsed.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id, ok := claims["sub"]
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user_id", id)

		c.Next()
	}
}

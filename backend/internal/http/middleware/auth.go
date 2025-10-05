package middleware

import (
	"FIXit/backend/internal/apierr"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
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
			AbortErr(c, apierr.Wrap(apierr.ErrUnauthorized, err))
			return
		}

		id, ok := claims["sub"]
		if !ok {
			AbortErr(c, apierr.Wrap(apierr.ErrUnauthorized, err))
			return
		}

		c.Set("user_id", id)

		c.Next()
	}
}

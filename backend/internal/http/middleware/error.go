package middleware

import (
	"FIXit/backend/internal/apierr"
	"github.com/gin-gonic/gin"
)

func Responder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		ae := apierr.From(err)

		// Пример: можно дополнительно приложить request-id, если он у вас где-то в контексте
		// rid, _ := c.Get("request_id")

		c.JSON(ae.HTTPStatus, gin.H{
			"error": gin.H{
				"code":    ae.Code,
				"message": ae.Message,
				"details": ae.Details,
				// "request_id": rid,
			},
		})
	}
}

func AbortErr(c *gin.Context, err error) {
	_ = c.Error(err)
	c.Abort()
}

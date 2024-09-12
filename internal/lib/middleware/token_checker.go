package middleware

import (
	"fmt"
	"net/http"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/jwt"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("access denied"),
			})
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		if claims.IP != c.ClientIP() {
			// TODO: send email message
		}
	}
}

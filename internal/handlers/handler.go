package handlers

import (
	"net/http"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/jwt"
	"github.com/gin-gonic/gin"
	log "github.com/rs/zerolog/log"
)

func GetTokensByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			GUID string `json:"guid" binding:"required"`
		}

		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Err(err)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error binding json",
			})

			return
		}

		accessToken, refreshToken, err := jwt.TokenGenerator(req.GUID)
		if err != nil {
			log.Err(err).Msg("error binding json")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":         accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func RefreshTokensByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			Token string `json:"token" binding:"required"`
		}

		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Err(err).Msg("error binding json")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

			return
		}

		claims, err := jwt.ValidateToken(req.Token)
		if err != nil {
			log.Err(err).Msg("error validating token")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

			return
		}

		if claims.Valid() != nil {
			log.Err(err)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

			return
		}

		accessToken, refreshToken, err := jwt.TokenGenerator(claims.Uid)

		c.JSON(http.StatusOK, gin.H{
			"token":         accessToken,
			"refresh_token": refreshToken,
		})
	}
}

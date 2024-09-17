package handlers

import (
	"errors"
	"net/http"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/lib/jwt"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/rs/zerolog/log"
)

var (
	ErrorRefresh   = errors.New("error refresh token")
	ErrorGetTokens = errors.New("error get tokens")
)

type UserHandler struct {
	up UserProvider
}

type UserProvider interface {
	GetUserByID(id int) (models.User, error)
	UpdateUserByID(id int, user models.User) error
}

func New(userProvider UserProvider) *UserHandler {
	return &UserHandler{
		up: userProvider,
	}
}

func (u *UserHandler) GetTokensByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			GUID int `json:"guid" binding:"required"`
		}

		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Err(err).Msg("error binding json")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrorGetTokens,
			})

			return
		}

		user, err := u.up.GetUserByID(req.GUID)
		if err != nil {
			log.Err(err).Msgf("error get user by id: %v", req.GUID)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrorGetTokens,
			})

			return
		}

		accessToken, refreshToken, err := jwt.TokenGenerator(req.GUID, c.ClientIP())
		if err != nil {
			log.Err(err).Msg("error generate token")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrorGetTokens,
			})

			return
		}

		user.RefreshToken = refreshToken
		err = u.up.UpdateUserByID(req.GUID, user)
		if err != nil {
			log.Err(err).Msgf("error update user by id: %v", req.GUID)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ErrorGetTokens,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":         accessToken,
			"refresh_token": refreshToken,
		})
	}
}

func (u *UserHandler) RefreshTokensByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		type Request struct {
			AccessToken  string `json:"token" binding:"required"`
			RefreshToken string `json:"refresh_token" binding:"required"`
		}

		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Err(err).Msg("error binding json")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrorRefresh,
			})

			return
		}

		claims, err := jwt.ValidateToken(req.AccessToken)
		if err != nil {
			log.Err(err).Msg("error validating token")

			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrorRefresh,
			})

			return
		}

		if err := claims.Valid(); err != nil {
			log.Err(err).Msgf("error require tokens %v user", claims.Uid)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrorRefresh,
			})

			return
		}

		err = jwt.RequireTokens(req.AccessToken, req.RefreshToken)
		if err != nil {
			log.Err(err).Msgf("error require tokens %v user", claims.Uid)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ErrorRefresh,
			})

			return
		}

		user, err := u.up.GetUserByID(claims.Uid)
		if err != nil {
			log.Err(err).Msgf("error get user by id: %v", claims.Uid)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrorGetTokens,
			})

			return
		}

		if user.RefreshToken != req.RefreshToken {
			log.Err(err).Msgf("error require tokens %v user", claims.Uid)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ErrorRefresh,
			})

			return
		}

		accessToken, refreshToken, err := jwt.TokenGenerator(claims.Uid, claims.IP)
		if err != nil {
			log.Err(err).Msgf("error generate jwt for %v user", claims.Uid)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ErrorRefresh,
			})

			return
		}

		user.RefreshToken = refreshToken
		err = u.up.UpdateUserByID(claims.Uid, user)
		if err != nil {
			log.Err(err).Msgf("error update user by id: %v", claims.Uid)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": ErrorRefresh,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":         accessToken,
			"refresh_token": refreshToken,
		})
	}
}

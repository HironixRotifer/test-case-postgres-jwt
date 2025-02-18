package auth

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/converter"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/http/request"
	"github.com/HironixRotifer/test-case-postgres-jwt/internal/service"
	"github.com/go-playground/validator/v10"
	log "github.com/rs/zerolog/log"
)

// @Summary Авторизация
// @Security Api-Token
// @Tags Auth
// @Description Авторизация
// @Produce json
// @Param authQuery body request.Credentials true "Модель авторизации"
// @Success 200
// @Failure 400 "Ошибки валидации входных параметров"
// @Failure 500 "Неизвестная внутренняя ошибка"
// @Router /api/v1/auth [post]
func Auth(jwtCustomService service.JWTCustomService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cr := &request.Credentials{}
		err := json.NewDecoder(r.Body).Decode(cr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error().Msgf("error: %v", err)
			return
		}

		uid := r.URL.Query().Get("uid")
		uidInt, err := strconv.Atoi(uid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error().Msgf("error %v", err)

			return
		}

		tokens, err := jwtCustomService.GetAccessToken(ctx, cr.Login, cr.Password, uidInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error().Msgf("error %v", err)
			return
		}

		err = json.NewEncoder(w).Encode(converter.ToResponseJwtCustomFromModel(tokens))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error().Msgf("error %v", err)

			return
		}
	}
}

// @Summary Обновление токенов
// @Security jwt
// @Tags Auth
// @Description Обновление токенов
// @Produce json
// @Param authQuery body request.JwtCustom true "Модель токенов"
// @Success 200 {object} response.JwtCustom
// @Failure 400 "Ошибки валидации входных параметров"
// @Failure 500 "Неизвестная внутренняя ошибка"
// @Router /api/v1/refresh [post]
func RefreshTokens(jwtCustomService service.JWTCustomService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jwtQuery request.JwtCustom

		ctx := r.Context()

		err := json.NewDecoder(r.Body).Decode(&jwtQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error().Msgf("error: %v", err)
			return
		}

		validate := validator.New()
		if err := validate.Struct(&jwtQuery); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error().Msgf("error: %v", err)
			return
		}

		tokens, err := jwtCustomService.RefreshToken(ctx, converter.ToModelJwtCustomFromRequest(&jwtQuery))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error().Msgf("error: %v", err)
			return
		}

		err = json.NewEncoder(w).Encode(converter.ToResponseJwtCustomFromModel(tokens))
		if err != nil {
			return
		}
	}
}

package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		models.NewErrorResponse(c, http.StatusUnauthorized, errors.ErrEmptyAuthHeader.Error())
		return
	}

	token, _ := strings.CutPrefix(header, "Bearer ")

	if err := h.service.ValidateToken(token); err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	login, err := parseToken(token)
	if err != nil {
		models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("token", token)
	c.Set("login", login)
}

func getLogin(c *gin.Context) (string, error) {
	login, ok := c.Get("login")
	if !ok {
		return "", errors.ErrLoginNotFound
	}

	loginString, ok := login.(string)
	if !ok {
		return "", errors.ErrInvalidLoginType
	}

	return loginString, nil
}

func parseToken(token string) (string, error) {
	token_, err := jwt.ParseWithClaims(token, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidSigningMethod
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token_.Claims.(*models.TokenClaims)
	if !ok {
		return "", errors.ErrInvalidTokenClaims
	}

	return claims.Login, nil
}

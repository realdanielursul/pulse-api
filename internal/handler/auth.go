package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (h *Handler) register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userProfile, err := h.service.Register(req)
	if err != nil {
		if err == errors.ErrInvalidLogin || err == errors.ErrInvalidEmail ||
			err == errors.ErrInvalidPassword || err == errors.ErrInvalidCountryCode ||
			err == errors.ErrInvalidPhone || err == errors.ErrInvalidImage {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err == errors.ErrLoginExists || err == errors.ErrPhoneExists {
			models.NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, userProfile)
}

func (h *Handler) signIn(c *gin.Context) {
	var req models.SignInRequest
	if err := c.BindJSON(&req); err != nil {
		models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.SignIn(req)
	if err != nil {
		if err == errors.ErrInvalidUsernameOrPassword {
			models.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/entity"
	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (h *Handler) register(c *gin.Context) {
	var req entity.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userProfile, err := h.service.Register(req)
	if err != nil {
		switch err {
		case errors.ErrInvalidLogin, errors.ErrInvalidEmail,
			errors.ErrInvalidPassword, errors.ErrInvalidCountryCode,
			errors.ErrInvalidPhone, errors.ErrInvalidImage:
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())

		case errors.ErrLoginExists, errors.ErrPhoneExists, errors.ErrEmailExists:
			entity.NewErrorResponse(c, http.StatusConflict, err.Error())

		default:
			entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	c.JSON(http.StatusCreated, userProfile)
}

func (h *Handler) signIn(c *gin.Context) {
	var req entity.SignInRequest
	if err := c.BindJSON(&req); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.SignIn(req)
	if err != nil {
		switch err {
		case errors.ErrInvalidUsernameOrPassword:
			entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())

		default:
			entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

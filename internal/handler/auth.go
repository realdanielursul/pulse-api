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
		if err == errors.ErrInvalidLogin || err == errors.ErrInvalidEmail ||
			err == errors.ErrInvalidPassword || err == errors.ErrInvalidCountryCode ||
			err == errors.ErrInvalidPhone || err == errors.ErrInvalidImage {
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err == errors.ErrLoginExists || err == errors.ErrPhoneExists {
			entity.NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
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
		if err == errors.ErrInvalidUsernameOrPassword {
			entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/entity"
	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (h *Handler) getProfile(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	userProfile, err := h.service.GetMyProfile(login)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func (h *Handler) updateProfile(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req entity.UpdateProfileRequest
	if err := c.BindJSON(&req); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userProfile, err := h.service.UpdateProfile(login, req)
	if err != nil {
		if err == errors.ErrInvalidCountryCode ||
			err == errors.ErrInvalidPhone ||
			err == errors.ErrInvalidImage {
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err == errors.ErrPhoneExists {
			entity.NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func (h *Handler) updatePassword(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req entity.UpdatePasswordRequest
	if err := c.BindJSON(&req); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.UpdatePassword(login, req); err != nil {
		if err == errors.ErrInvalidPassword {
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err == errors.ErrInvalidOldPassword {
			entity.NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) getProfileByLogin(c *gin.Context) {
	userLogin, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	profileLogin := c.Param("login")
	if profileLogin == "" {
		entity.NewErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidLogin.Error())
		return
	}

	userProfile, err := h.service.GetProfileByLogin(userLogin, profileLogin)
	if err != nil {
		if err == errors.ErrLoginDoesNotExist || err == errors.ErrAccessDenied {
			entity.NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realdanielursul/pulse-api/internal/service"
)

func (h *Handler) getMyProfile(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetMyProfile(c.Request.Context(), login)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

type updateProfileInput struct {
	CountryCode *string `json:"countryCode"`
	IsPublic    *bool   `json:"isPublic"`
	Phone       *string `json:"phone"`
	Image       *string `json:"image"`
}

func (h *Handler) updateProfile(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input := &updateProfileInput{}
	if err := c.BindJSON(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.UpdateProfile(c.Request.Context(), login, &service.UserUpdateProfileInput{
		CountryCode: input.CountryCode,
		IsPublic:    input.IsPublic,
		Phone:       input.Phone,
		Image:       input.Image,
	})
	if err != nil {
		if err == service.ErrCountryNotFound {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err == service.ErrPhoneAlreadyExists {
			NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getProfile(c *gin.Context) {
	login := c.Param("login")
	if login == "" {
		NewErrorResponse(c, http.StatusBadRequest, "empty login param")
		return
	}

	requesterLogin, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.GetProfile(c.Request.Context(), login, requesterLogin)
	if err != nil {
		if err == service.ErrAccessDenied || err == service.ErrUserNotFound {
			NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

type updatePasswordInput struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func (h *Handler) updatePassword(c *gin.Context) {
	input := &updatePasswordInput{}
	if err := c.BindJSON(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Auth.UpdatePassword(c.Request.Context(), login, &service.AuthUpdatePasswordInput{
		OldPassword: input.OldPassword,
		NewPassword: input.NewPassword,
	})
	if err != nil {
		if err == service.ErrInvalidLoginOrPassword {
			NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

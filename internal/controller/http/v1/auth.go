package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realdanielursul/pulse-api/internal/service"
)

type registerInput struct {
	Login       string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CountryCode string `json:"countryCode"`
	IsPublic    bool   `json:"isPublic"`
	Phone       string `json:"phone"`
	Image       string `json:"image"`
}

func (h *Handler) register(c *gin.Context) {
	input := &registerInput{}
	if err := c.BindJSON(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Auth.Register(c.Request.Context(), &service.AuthRegisterInput{
		Login:       input.Login,
		Email:       input.Email,
		Password:    input.Password,
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

		if err == service.ErrLoginAlreadyExists || err == service.ErrEmailAlreadyExists || err == service.ErrPhoneAlreadyExists {
			NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

type signInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	input := &signInInput{}
	if err := c.BindJSON(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Auth.SignIn(c.Request.Context(), &service.AuthSignInInput{
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		if err == service.ErrInvalidLoginOrPassword || err == service.ErrCannotSignToken {
			NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

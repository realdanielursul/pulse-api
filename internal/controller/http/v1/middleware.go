package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	token, _ := strings.CutPrefix(header, "Bearer ")

	login, isValid, err := h.services.Auth.ValidateToken(c.Request.Context(), token)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !isValid {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid token")
		return
	}

	c.Set("token", token)
	c.Set("login", login)
}

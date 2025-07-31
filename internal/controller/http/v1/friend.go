package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realdanielursul/pulse-api/internal/service"
)

type addFriendInput struct {
	Login string `json:"login"`
}

func (h *Handler) addFriend(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input := &addFriendInput{}
	if err := c.BindJSON(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if login == input.Login {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	err = h.services.Friend.AddFriend(c.Request.Context(), input.Login, login)
	if err != nil {
		if err == service.ErrUserNotFound {
			NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

type removeFriendInput struct {
	Login string `json:"login"`
}

func (h *Handler) removeFriend(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input := &removeFriendInput{}
	if err := c.BindJSON(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if login == input.Login {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	err = h.services.Friend.RemoveFriend(c.Request.Context(), input.Login, login)
	if err != nil {
		if err == service.ErrUserNotFound {
			NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) listFriends(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	friends, err := h.services.Friend.ListFriends(c.Request.Context(), login, limit, offset)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, friends)
}

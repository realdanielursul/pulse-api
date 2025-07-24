package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/entity"
	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (h *Handler) addFriend(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req entity.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.AddFriend(login, req.Login); err != nil {
		if err == errors.ErrLoginDoesNotExist {
			entity.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) removeFriend(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req entity.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.RemoveFriend(login, req.Login); err != nil {
		if err == errors.ErrLoginDoesNotExist {
			entity.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *Handler) listFriends(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	limit_ := c.Query("limit")
	offset_ := c.Query("offset")

	limit, err := strconv.Atoi(limit_)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	offset, err := strconv.Atoi(offset_)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	friends, err := h.service.ListFriends(login, limit, offset)
	if err != nil {
		if err == errors.ErrInvalidPaginationParams {
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, friends)
}

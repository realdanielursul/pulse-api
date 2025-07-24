package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ursulgwopp/pulse-api/internal/entity"
	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (h *Handler) newPost(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var req entity.NewPostRequest
	if err := c.BindJSON(&req); err != nil {
		entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.service.NewPost(login, req)
	if err != nil {
		if err == errors.ErrInvalidTag || err == errors.ErrInvalidContent {
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) getPost(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	postIdParam := c.Param("postId")
	if postIdParam == "" {
		entity.NewErrorResponse(c, http.StatusNotFound, errors.ErrPostIdNotFound.Error())
		return
	}

	postId, err := uuid.Parse(postIdParam)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusNotFound, errors.ErrPostIdNotFound.Error())
		return
	}

	post, err := h.service.GetPost(login, postId)
	if err != nil {
		if err == errors.ErrPostIdNotFound || err == errors.ErrAccessDenied {
			entity.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) listMyPosts(c *gin.Context) {
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

	posts, err := h.service.ListMyPosts(login, limit, offset)
	if err != nil {
		if err == errors.ErrInvalidPaginationParams {
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) listPosts(c *gin.Context) {
	userLogin, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	login := c.Param("login")
	if login == "" {
		entity.NewErrorResponse(c, http.StatusBadRequest, errors.ErrLoginNotFound.Error())
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

	posts, err := h.service.ListPosts(userLogin, login, limit, offset)
	if err != nil {
		if err == errors.ErrInvalidPaginationParams {
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err == errors.ErrAccessDenied || err == errors.ErrLoginDoesNotExist {
			entity.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) likePost(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	postIdParam := c.Param("postId")
	if postIdParam == "" {
		entity.NewErrorResponse(c, http.StatusUnauthorized, errors.ErrPostIdNotFound.Error())
		return
	}

	postId, err := uuid.Parse(postIdParam)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, errors.ErrPostIdNotFound.Error())
		return
	}

	post, err := h.service.LikePost(login, postId)
	if err != nil {
		if err == errors.ErrAccessDenied || err == errors.ErrLoginDoesNotExist {
			entity.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) dislikePost(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	postIdParam := c.Param("postId")
	if postIdParam == "" {
		entity.NewErrorResponse(c, http.StatusUnauthorized, errors.ErrPostIdNotFound.Error())
		return
	}

	postId, err := uuid.Parse(postIdParam)
	if err != nil {
		entity.NewErrorResponse(c, http.StatusUnauthorized, errors.ErrPostIdNotFound.Error())
		return
	}

	post, err := h.service.DislikePost(login, postId)
	if err != nil {
		if err == errors.ErrAccessDenied || err == errors.ErrLoginDoesNotExist {
			entity.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

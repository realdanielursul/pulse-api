package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realdanielursul/pulse-api/internal/service"
)

type createPostInput struct {
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (h *Handler) createPost(c *gin.Context) {
	input := &createPostInput{}
	if err := c.BindJSON(input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.services.Post.CreatePost(c.Request.Context(), &service.PostCreatePostInput{
		Content: input.Content,
		Author:  login,
		Tags:    input.Tags,
	})
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) getPost(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		NewErrorResponse(c, http.StatusBadRequest, "error post id empty")
		return
	}

	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := h.services.Post.GetPost(c.Request.Context(), postId, login)
	if err != nil {
		if err == service.ErrPostNotFound || err == service.ErrAccessDenied {
			NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) getMyPosts(c *gin.Context) {
	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
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

	posts, err := h.services.Post.GetMyFeed(c.Request.Context(), login, limit, offset)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) getUserPosts(c *gin.Context) {
	userLogin := c.Param("login")
	if userLogin == "" {
		NewErrorResponse(c, http.StatusInternalServerError, "empty login param")
		return
	}

	requesterLogin, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
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

	posts, err := h.services.Post.GetUserFeed(c.Request.Context(), userLogin, requesterLogin, limit, offset)
	if err != nil {
		if err == service.ErrAccessDenied {
			NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *Handler) likePost(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		NewErrorResponse(c, http.StatusBadRequest, "empty post id param")
		return
	}

	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.services.Post.LikePost(c.Request.Context(), postId, login)
	if err != nil {
		if err == service.ErrPostNotFound {
			NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if err == service.ErrAccessDenied {
			NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) dislikePost(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		NewErrorResponse(c, http.StatusBadRequest, "empty post id param")
		return
	}

	login, err := getLogin(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.services.Post.DislikePost(c.Request.Context(), postId, login)
	if err != nil {
		if err == service.ErrPostNotFound {
			NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if err == service.ErrAccessDenied {
			NewErrorResponse(c, http.StatusForbidden, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ursulgwopp/pulse-api/internal/entity"
)

type Service interface {
	ListCountries(regions []string) ([]entity.Country, error)
	GetCountryByAlpha2(alpha2 string) (entity.Country, error)

	Register(req entity.RegisterRequest) (entity.UserProfile, error)
	SignIn(req entity.SignInRequest) (string, error)
	ValidateToken(token string) error

	GetMyProfile(login string) (entity.UserProfile, error)
	UpdateProfile(login string, req entity.UpdateProfileRequest) (entity.UserProfile, error)
	UpdatePassword(login string, req entity.UpdatePasswordRequest) error
	GetProfileByLogin(userLogin string, profileLogin string) (entity.UserProfile, error)

	AddFriend(userLogin string, login string) error
	RemoveFriend(userLogin string, login string) error
	ListFriends(login string, limit int, offset int) ([]entity.FriendInfo, error)

	NewPost(login string, req entity.NewPostRequest) (entity.Post, error)
	GetPost(login string, postId uuid.UUID) (entity.Post, error)
	ListMyPosts(login string, limit int, offset int) ([]entity.Post, error)
	ListPosts(userLogin string, login string, limit int, offset int) ([]entity.Post, error)
	LikePost(login string, postId uuid.UUID) (entity.Post, error)
	DislikePost(login string, postId uuid.UUID) (entity.Post, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Serve the OpenAPI YAML file
	router.GET("/openapi.yml", func(c *gin.Context) {
		c.File("openapi.yml") // Adjust the path as necessary
	})

	// Serve Swagger UI directly
	router.GET("/swagger/", func(c *gin.Context) {
		// Serve the Swagger UI HTML directly
		c.Header("Content-Type", "text/html")
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		 <meta charset="utf-8" />
		 <meta name="viewport" content="width=device-width, initial-scale=1" />
		 <meta name="description" content="SwaggerUI" />
		 <title>SwaggerUI</title>
		 <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
		</head>
		<body>
		<div id="swagger-ui"></div>
		<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
		<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-standalone-preset.js" crossorigin></script>
		<script>
		 window.onload = () => {
		 window.ui = SwaggerUIBundle({
		  url: 'http://localhost:8080/openapi.yml',
		  dom_id: '#swagger-ui',
		  presets: [
		  SwaggerUIBundle.presets.apis,
		  SwaggerUIStandalonePreset
		  ],
		  layout: "StandaloneLayout",
		 });
		 };
		</script>
		</body>
		</html>
			  `
		c.String(http.StatusOK, html)
	})

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})

		api.GET("/countries", h.listCountries)
		api.GET("/countries/:alpha2", h.getCountryByAlpha2)

		api.POST("/auth/register", h.register)
		api.POST("/auth/sign-in", h.signIn)

		api.GET("/me/profile", h.userIdentity, h.getProfile)
		api.PATCH("/me/profile", h.userIdentity, h.updateProfile)
		api.POST("/me/updatePassword", h.userIdentity, h.updatePassword)

		api.GET("/profiles/:login", h.userIdentity, h.getProfileByLogin)

		api.POST("/friends/add", h.userIdentity, h.addFriend)
		api.POST("/friends/remove", h.userIdentity, h.removeFriend)
		api.GET("/friends", h.userIdentity, h.listFriends)

		api.POST("/posts/new", h.userIdentity, h.newPost)
		api.GET("/posts/:postId", h.userIdentity, h.getPost)
		api.GET("/posts/feed/my", h.userIdentity, h.listMyPosts)
		api.GET("/posts/feed/:login", h.userIdentity, h.listPosts)
		api.POST("/posts/:postId/like", h.userIdentity, h.likePost)
		api.POST("/posts/:postId/dislike", h.userIdentity, h.dislikePost)
	}

	return router
}

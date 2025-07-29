package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realdanielursul/pulse-api/internal/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
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
			c.String(http.StatusOK, "ok")
		})

		api.GET("/countries", h.listCountries)
		api.GET("/countries/:alpha2", h.getCountry)
	}

	return router
}

type ErrorResponse struct {
	Reason interface{} `json:"reason"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Reason: message})
}

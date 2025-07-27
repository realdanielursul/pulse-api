package v1

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
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

	return router
}

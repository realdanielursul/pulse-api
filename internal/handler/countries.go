package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (h *Handler) listCountries(c *gin.Context) {
	regions := c.QueryArray("region")

	countries, err := h.service.ListCountries(regions)
	if err != nil {
		if err == errors.ErrInvalidRegion {
			models.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, countries)
}

func (h *Handler) getCountryByAlpha2(c *gin.Context) {
	alpha2 := c.Param("alpha2")
	if alpha2 == "" {
		models.NewErrorResponse(c, http.StatusBadRequest, errors.ErrEmptyAlpha2.Error())
		return
	}

	country, err := h.service.GetCountryByAlpha2(alpha2)
	if err != nil {
		if err == errors.ErrCountryNotFound {
			models.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		models.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, country)
}

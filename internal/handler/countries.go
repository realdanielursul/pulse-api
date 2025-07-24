package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ursulgwopp/pulse-api/internal/entity"
	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (h *Handler) listCountries(c *gin.Context) {
	regions := c.QueryArray("region")

	countries, err := h.service.ListCountries(regions)
	if err != nil {
		if err == errors.ErrInvalidRegion {
			entity.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, countries)
}

func (h *Handler) getCountryByAlpha2(c *gin.Context) {
	alpha2 := c.Param("alpha2")
	if alpha2 == "" {
		entity.NewErrorResponse(c, http.StatusBadRequest, errors.ErrEmptyAlpha2.Error())
		return
	}

	country, err := h.service.GetCountryByAlpha2(alpha2)
	if err != nil {
		if err == errors.ErrCountryNotFound {
			entity.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		entity.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, country)
}

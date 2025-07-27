package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realdanielursul/pulse-api/internal/model"
	"github.com/realdanielursul/pulse-api/internal/service"
)

func (h *Handler) listCountries(c *gin.Context) {
	regions := c.QueryArray("region")

	countries, err := h.service.ListCountries(regions)
	if err != nil {
		if err == service.ErrInvalidRegion {
			model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, countries)
}

func (h *Handler) getCountryByAlpha2(c *gin.Context) {
	alpha2 := c.Param("alpha2")
	if alpha2 == "" {
		model.NewErrorResponse(c, http.StatusBadRequest, "empty alpha2 code")
		return
	}

	country, err := h.service.GetCountryByAlpha2(alpha2)
	if err != nil {
		if err == service.ErrCountryNotFound {
			model.NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, country)
}

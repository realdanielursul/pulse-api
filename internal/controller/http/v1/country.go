package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realdanielursul/pulse-api/internal/service"
)

func (h *Handler) listCountries(c *gin.Context) {
	regions := c.QueryArray("region")

	countries, err := h.services.Country.ListCountries(c.Request.Context(), regions)
	if err != nil {
		if err == service.ErrInvalidRegion {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, countries)
}

func (h *Handler) getCountry(c *gin.Context) {
	alpha2 := c.Param("alpha2")
	if alpha2 == "" {
		NewErrorResponse(c, http.StatusBadRequest, "empty alpha2 code")
		return
	}

	country, err := h.services.Country.GetCountry(c.Request.Context(), alpha2)
	if err != nil {
		if err == service.ErrCountryNotFound {
			NewErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, country)
}

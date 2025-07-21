package service

import (
	"sort"

	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (s *Service) ListCountries(regions []string) ([]models.Country, error) {
	if regions[0] != "" && len(regions) != 0 {
		for _, region := range regions {
			if !isValidRegion(region) {
				return []models.Country{}, errors.ErrInvalidRegion
			}
		}
	}

	countries, err := s.repo.ListCountries(regions)
	if err != nil {
		return []models.Country{}, err
	}

	sortByRegion := func(i, j int) bool {
		return countries[i].Alpha2 < countries[j].Alpha2
	}
	sort.Slice(countries, sortByRegion)

	return countries, nil
}

func (s *Service) GetCountryByAlpha2(alpha2 string) (models.Country, error) {
	country, err := s.repo.GetCountryByAlpha2(alpha2)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return models.Country{}, errors.ErrCountryNotFound
		}

		return models.Country{}, err
	}

	return country, nil
}

package service

import (
	"database/sql"
	"sort"

	"github.com/ursulgwopp/pulse-api/internal/errors"
	"github.com/ursulgwopp/pulse-api/internal/models"
)

func (s *Service) ListCountries(regions []string) ([]models.Country, error) {
	if len(regions) != 0 {
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

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Alpha2 < countries[j].Alpha2
	})

	return countries, nil
}

func (s *Service) GetCountryByAlpha2(alpha2 string) (models.Country, error) {
	country, err := s.repo.GetCountryByAlpha2(alpha2)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Country{}, errors.ErrCountryNotFound
		}

		return models.Country{}, err
	}

	return country, nil
}

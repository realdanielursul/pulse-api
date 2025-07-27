package service

import (
	"database/sql"
	"slices"
	"sort"

	"github.com/realdanielursul/pulse-api/internal/model"
)

func (s *Service) ListCountries(regions []string) ([]model.Country, error) {
	if len(regions) != 0 {
		for _, region := range regions {
			if !isValidRegion(region) {
				return []model.Country{}, ErrInvalidRegion
			}
		}
	}

	countries, err := s.repo.ListCountries(regions)
	if err != nil {
		return []model.Country{}, err
	}

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Alpha2 < countries[j].Alpha2
	})

	return countries, nil
}

func (s *Service) GetCountryByAlpha2(alpha2 string) (model.Country, error) {
	country, err := s.repo.GetCountryByAlpha2(alpha2)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Country{}, ErrCountryNotFound
		}

		return model.Country{}, err
	}

	return country, nil
}

func isValidRegion(region string) bool {
	validRegions := []string{"Asia", "Oceania", "Europe", "Africa", "Americas"}
	return slices.Contains(validRegions, region)
}

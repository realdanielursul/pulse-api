package service

import (
	"database/sql"
	"sort"

	"github.com/ursulgwopp/pulse-api/internal/entity"
	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (s *Service) ListCountries(regions []string) ([]entity.Country, error) {
	if len(regions) != 0 {
		for _, region := range regions {
			if !isValidRegion(region) {
				return []entity.Country{}, errors.ErrInvalidRegion
			}
		}
	}

	countries, err := s.repo.ListCountries(regions)
	if err != nil {
		return []entity.Country{}, err
	}

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Alpha2 < countries[j].Alpha2
	})

	return countries, nil
}

func (s *Service) GetCountryByAlpha2(alpha2 string) (entity.Country, error) {
	country, err := s.repo.GetCountryByAlpha2(alpha2)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Country{}, errors.ErrCountryNotFound
		}

		return entity.Country{}, err
	}

	return country, nil
}

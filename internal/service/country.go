package service

import (
	"context"
	"slices"
	"sort"

	"github.com/realdanielursul/pulse-api/internal/entity"
	"github.com/realdanielursul/pulse-api/internal/repository"
)

type CountryService struct {
	countryRepo repository.Country
}

func NewCountryService(countryRepo repository.Country) *CountryService {
	return &CountryService{
		countryRepo: countryRepo,
	}
}

func (s *CountryService) ListCountries(ctx context.Context, regions []string) ([]*CountryOutput, error) {
	countries := make([]*entity.Country, 0, 100)
	var err error

	if len(regions) > 0 {
		for _, region := range regions {
			if !isValidRegion(region) {
				return nil, ErrInvalidRegion
			}
		}

		countries, err = s.countryRepo.GetCountriesByRegion(ctx, regions)
	} else {
		countries, err = s.countryRepo.GetAllCountries(ctx)
	}

	if err != nil {
		return nil, err
	}

	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Alpha2 < countries[j].Alpha2
	})

	outputCountries := make([]*CountryOutput, 0, len(countries))
	for _, country := range countries {
		outputCountry := &CountryOutput{
			Name:   country.Region,
			Alpha2: country.Alpha2,
			Alpha3: country.Alpha2,
			Region: country.Region,
		}

		outputCountries = append(outputCountries, outputCountry)
	}

	return outputCountries, nil
}

func (s *CountryService) GetCountry(ctx context.Context, alpha2 string) (*CountryOutput, error) {
	country, err := s.countryRepo.GetCountryByAlpha2(ctx, alpha2)
	if err != nil {
		if country == nil {
			return nil, ErrCountryNotFound
		}

		return nil, err
	}

	return &CountryOutput{
		Name:   country.Name,
		Alpha2: country.Alpha2,
		Alpha3: country.Alpha3,
		Region: country.Region,
	}, nil
}

func isValidRegion(region string) bool {
	validRegions := []string{"Asia", "Oceania", "Europe", "Africa", "Americas"}
	return slices.Contains(validRegions, region)
}

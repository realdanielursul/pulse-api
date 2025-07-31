package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/realdanielursul/pulse-api/internal/entity"
)

type CountryRepository struct {
	*sqlx.DB
}

func NewCountryRepository(db *sqlx.DB) *CountryRepository {
	return &CountryRepository{db}
}

func (r *CountryRepository) GetAllCountries(ctx context.Context) ([]*entity.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	countries := make([]*entity.Country, 0, 100)
	sql := `SELECT name, alpha2, alpha3, region FROM countries`
	rows, err := r.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		country := &entity.Country{}
		if err := rows.Scan(&country.Name, &country.Alpha2, &country.Alpha3, &country.Region); err != nil {
			return nil, err
		}

		countries = append(countries, country)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return countries, nil
}

func (r *CountryRepository) GetCountriesByRegion(ctx context.Context, regions []string) ([]*entity.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	countries := make([]*entity.Country, 0, 100)
	for _, region := range regions {
		query := `SELECT name, alpha2, alpha3, region FROM countries WHERE region = $1`
		rows, err := r.QueryContext(ctx, query, region)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			country := &entity.Country{}
			if err := rows.Scan(&country.Name, &country.Alpha2, &country.Alpha3, &country.Region); err != nil {
				return nil, err
			}

			countries = append(countries, country)
		}

		if rows.Err() != nil {
			return nil, err
		}
	}

	return countries, nil
}

func (r *CountryRepository) GetCountryByAlpha2(ctx context.Context, alpha2 string) (*entity.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	country := &entity.Country{}
	query := `SELECT name, alpha2, alpha3, region FROM countries WHERE alpha2 = $1`
	if err := r.QueryRowContext(ctx, query, alpha2).Scan(&country.Name, &country.Alpha2, &country.Alpha3, &country.Region); err != nil {
		return nil, err
	}

	return country, nil
}

package repository

import (
	"context"
	"database/sql"

	"github.com/ursulgwopp/pulse-api/internal/entity"
)

func (r *PostgresRepository) ListCountries(regions []string) ([]entity.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var countries []entity.Country
	var rows *sql.Rows
	var err error

	if len(regions) == 0 {
		query := `SELECT name, alpha2, alpha3, region FROM countries`
		rows, err = r.db.QueryContext(ctx, query)

		if err != nil {
			return []entity.Country{}, err
		}

		for rows.Next() {
			var country entity.Country
			if err := rows.Scan(&country.Name, &country.Alpha2, &country.Alpha3, &country.Region); err != nil {
				return []entity.Country{}, err
			}

			countries = append(countries, country)
		}

		if rows.Err() != nil {
			return []entity.Country{}, err
		}

		return countries, nil
	}

	for _, region := range regions {
		query := `SELECT name, alpha2, alpha3, region FROM countries WHERE region = $1`
		rows, err = r.db.QueryContext(ctx, query, region)

		if err != nil {
			return []entity.Country{}, err
		}

		for rows.Next() {
			var country entity.Country
			if err := rows.Scan(&country.Name, &country.Alpha2, &country.Alpha3, &country.Region); err != nil {
				return []entity.Country{}, err
			}

			countries = append(countries, country)
		}

		if rows.Err() != nil {
			return []entity.Country{}, err
		}
	}

	return countries, nil
}

func (r *PostgresRepository) GetCountryByAlpha2(alpha2 string) (entity.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	var country entity.Country

	query := `SELECT name, alpha2, alpha3, region FROM countries WHERE alpha2 = $1`
	if err := r.db.QueryRowContext(ctx, query, alpha2).Scan(&country.Name, &country.Alpha2, &country.Alpha3, &country.Region); err != nil {
		return entity.Country{}, err
	}

	return country, nil
}

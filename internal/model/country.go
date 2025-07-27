package model

type Country struct {
	Name   string `json:"name" db:"name"`
	Alpha2 string `json:"alpha2" db:"alpha2"`
	Alpha3 string `json:"alpha3" db:"alpha3"`
	Region string `json:"region" db:"region"`
}

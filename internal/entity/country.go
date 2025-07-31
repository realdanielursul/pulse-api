package entity

type Country struct {
	Name   string `db:"name"`
	Alpha2 string `db:"alpha2"`
	Alpha3 string `db:"alpha3"`
	Region string `db:"region"`
}

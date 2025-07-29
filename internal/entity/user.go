package entity

type User struct {
	Login        string `db:"login"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
	CountryCode  string `db:"country_code"`
	IsPublic     bool   `db:"is_public"`
	Phone        string `db:"phone"`
	Image        string `db:"image"`
}

package entity

type Token struct {
	Login       string `db:"login"`
	TokenString string `db:"token_string"`
	IsValid     bool   `db:"is_valid"`
}

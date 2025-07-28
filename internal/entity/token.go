package entity

type Token struct {
	Login   string `db:"login"`
	Token   string `db:"token"`
	IsValid bool   `db:"is_valid"`
}

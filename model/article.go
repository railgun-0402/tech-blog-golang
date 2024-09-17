package model

type Article struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}

package repository

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func SetDB(d *sqlx.DB) {
	db = d
}

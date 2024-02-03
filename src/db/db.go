package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var Database *sqlx.DB

func InitDB() {
	db, err := sqlx.Open("sqlite3", "file:application.db?cache=shared")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1)
	Database = db
}

func Cleanup() {
	Database.Close()
}

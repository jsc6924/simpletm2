package config

import (
	"database/sql"
	"log"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("sqlite3", DBPath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	boil.SetDB(db)
	DB = db
}

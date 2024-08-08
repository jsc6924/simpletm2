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
	// Enable foreign key support
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	DB = db
}

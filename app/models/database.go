package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)



var db *sql.DB

func init()  {
	var err error
	db, err = sql.Open("postgres", "user=ryohei dbname=athlete_app password=38ryohei38 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

}
	

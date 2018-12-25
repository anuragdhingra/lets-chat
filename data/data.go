package data

import "database/sql"

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "dbname=chitchat sslmode=disable")
	if err!= nil {
		// Log the error
	}
	return
}

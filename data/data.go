package data

import (
	"database/sql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"


var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:password@tcp(localhost:3307)/chitchat?parseTime=true")
	if err!= nil {
		log.Print(err)
		return
	} else
	{
		log.Print("Successfully connected to db")
	}
	return
}
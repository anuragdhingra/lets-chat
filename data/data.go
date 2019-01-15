package data

import (
	"database/sql"
	"log"
	"os"
)
import _ "github.com/go-sql-driver/mysql"


var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", os.Getenv("datasource"))
	if err!= nil {
		log.Print(err)
		return
	} else
	{
		log.Print("Successfully connected to db")
	}
	return
}
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
	Db, err = sql.Open("mysql", getDatasource())
	if err!= nil {
		log.Print(err)
		return
	} else
	{
		log.Print("Successfully connected to datasource: " + getDatasource())
	}
	return
}

func getDatasource() (dataSource string) {
	dataSource = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_URL") +
		")/" + os.Getenv("DB_NAME") + "?parseTime=true"
	return
}


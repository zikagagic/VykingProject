package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMYSQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "vyking_project_db"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?multiStatements=true")
	return
}

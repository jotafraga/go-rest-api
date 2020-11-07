package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectDataBase() *sql.DB {
	//connection := "user=user dbname=tcc2020 password=userpw host=localhost sslmode=disable"
	connection := "user:userpw@/tcc2020"
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

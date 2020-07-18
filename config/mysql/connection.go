package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var(
	session *sql.DB
)

func InitSession() {
	dbConnection, err := sql.Open("mysql", "root:root@tcp(db:3306)/enchainte_db")

	if err != nil {
		panic(err)
	}
	if err = dbConnection.Ping(); err != nil {
		panic(fmt.Sprintf("error when trying to ping to the provided DSN: %s", err))
	}

	session = dbConnection
	fmt.Println("mysql session successfully configured")
}

func Session() *sql.DB {
	return session
}
package mysql

import (
	"database/sql"
	"fmt"
	"github.com/AlexRipoll/enchante_technical_interview/config"
	_ "github.com/go-sql-driver/mysql"
)

var(
	session *sql.DB
)

func InitSession() {
	dbConnection, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.Params.Database.User,
		config.Params.Database.Password,
		config.Params.Database.Host,
		config.Params.Database.Port,
		config.Params.Database.Name,
	))
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
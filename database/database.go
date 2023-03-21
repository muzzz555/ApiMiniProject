package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

const (
	host     = "root"
	// port = "localhost:3306"
	port = "localhost:3307"
	password = "1234"
	dbname   = "seproject"
)

func Connectdata() *sql.DB {
	fmt.Println("Connecting database")
	sqlinfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", host, password, port, dbname)
	db, err := sql.Open("mysql", sqlinfo)

	if err != nil {
		panic(err.Error())
	}
	return db
}

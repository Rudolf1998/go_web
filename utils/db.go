package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:1998@tcp(localhost:3306)/mydb?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

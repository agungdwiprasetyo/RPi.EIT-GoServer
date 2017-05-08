// modul yg berhubungan dgn database
package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DB_Connect() (*sql.DB,error) {
	db, err := sql.Open("mysql", "root:fakepasswordSS@tcp(localhost:3306)/eit")
	return db, err
}
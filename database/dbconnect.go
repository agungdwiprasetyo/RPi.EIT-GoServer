// modul yg berhubungan dgn database
package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"../models"
)

func Connect() {
	db, err := sql.Open("mysql", "root:fakepassword@tcp(localhost:3306)/eit")
	if err != nil {
		fmt.Print(err.Error())
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	// init database to models
	models.SetDB(db)
}
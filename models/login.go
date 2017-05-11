package models

import (
	"fmt"
)

func AuthLogin(user string) Login{
	var (
		login Login
	)

	rows, err := db.Query("select * from login where username=?",user)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&login.ID, &login.Tipe, &login.IDAlat, &login.Username, &login.Password, &login.Nama)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return login
}
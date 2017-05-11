package models

import (
	"database/sql"
)

type (
	Image struct {
		Nama string `json:"nama,omitempty"`
		Id_data int `json:"id_data,omitempty"`
		Id_algor string `json:"id_algor,omitempty"`
		Kerapatan float64 `json:"kerapatan,omitempty"`
		Datetime string `json:"datetime,omitempty"`
	}

	Data struct {
		Id_data int `json:"id_data,omitempty"`
		Nama string `json:"nama_data,omitempty"`
		Filename string `json:"filename,omitempty"`
		Deskripsi string `json:"deskripsi,omitempty"`
		Model string `json:"model,omitempty"`
		Citra string `json:"citra,omitempty"`
		Arus float64 `json:"arus,omitempty"`
		Datetime string `json:"datetime,omitempty"`
	}

	Algor struct {
		Id_algor string `json:"id_algor,omitempty"`
		Nama string `json:"nama,omitempty"`
	}

	Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
)

var db *sql.DB

func SetDB(conn *sql.DB){
	db = conn
}
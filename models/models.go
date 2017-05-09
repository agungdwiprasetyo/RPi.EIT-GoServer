package models

import (
	"database/sql"
	// "github.com/jinzhu/gorm"
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
)

var conn *sql.DB
func SetDB(db *sql.DB){
	conn = db
}
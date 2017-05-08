package models

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context){
	var (
		data  Data
		alldata []Data
	)
	rows, err := conn.Query("select * from data_ukur;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&data.Id_data, &data.Nama, &data.Filename, &data.Deskripsi, &data.Model, &data.Citra, &data.Arus, &data.Datetime)
		alldata = append(alldata, data)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"result": alldata,
		"count":  len(alldata),
	})
}

func PostData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "tes post data ok gan",
	})
}

func PutData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "tes put data ok gan",
	})
}

func DeleteData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "tes delete data ok gan",
	})
}
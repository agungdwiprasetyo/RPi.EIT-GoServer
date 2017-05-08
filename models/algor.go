package models

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAlgor(c *gin.Context){
	var (
		algor  Algor
		allAlgor []Algor
	)
	rows, err := conn.Query("select id_algor,nama_algor from algoritma;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&algor.Id_algor, &algor.Nama)
		allAlgor = append(allAlgor, algor)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": allAlgor,
		"count":  len(allAlgor),
	})
}

func PostAlgor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "tes post algor ok gan",
	})
}

func PutAlgor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "tes put algor ok gan",
	})
}

func DeleteAlgor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "tes delete algor ok gan",
	})
}
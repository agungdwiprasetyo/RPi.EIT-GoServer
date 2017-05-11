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

	rows, err := db.Query("select id_algor,nama_algor from algoritma;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&algor.Id_algor, &algor.Nama)
		allAlgor = append(allAlgor, algor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ops, ada kesalahan saat fetch data",
			})
			fmt.Print(err.Error())
			return
		}
	}
	defer rows.Close()

	c.JSON(http.StatusOK, allAlgor)
}

func PostAlgor(c *gin.Context) {
	namaAlgor := c.PostForm("nama_algor")
	idAlgor := c.PostForm("id_algor")

	if namaAlgor == "" || idAlgor == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "tidak ada parameter yang dikirim",
		})
		return
	}

	masuk, err := db.Prepare("insert into algoritma (id_algor,nama_algor) values(?,?);")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	_, err = masuk.Exec(idAlgor, namaAlgor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat eksekusi query",
		})
		fmt.Print(err.Error())
		return
	}
	defer masuk.Close()

	c.JSON(http.StatusCreated, gin.H{
		"message": "sukses tambah jenis algoritma baru",
	})
}

func PutAlgor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "tes put algor ok gan",
	})
}

func DeleteAlgor(c *gin.Context) {
	idAlgor := c.Param("id")
	if idAlgor == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "tidak ada parameter yang dikirim",
		})
		return
	}

	dlt, err := db.Prepare("delete from algoritma where id_algor=?;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	_, err = dlt.Exec(idAlgor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat eksekusi query",
		})
		fmt.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete jenis algoritma sukses gan",
	})
}
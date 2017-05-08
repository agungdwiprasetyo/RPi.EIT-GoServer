package models

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context){
	var (
		image  Image
		images []Image
	)
	rows, err := conn.Query("select * from image;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&image.Nama, &image.Id_data, &image.Id_algor, &image.Kerapatan, &image.Datetime)
		images = append(images, image)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": images,
		"count":  len(images),
	})
}

func PostImage(c *gin.Context) {
	nama := c.PostForm("filename")
	idData := c.PostForm("id_data")
	idAlgor := c.PostForm("id_algor")
	kerapatan := c.PostForm("kerapatan")

	masuk, err := conn.Prepare("insert into image (nama, id_data, id_algor, kerapatan) values(?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = masuk.Exec(nama, idData, idAlgor, kerapatan)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
		})
		fmt.Print(err.Error())
	}
	defer masuk.Close()

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "sukses tambah gambar",
	})
}

func PutImage(c *gin.Context) {
	nama := c.PostForm("filename")
	kerapatan := c.PostForm("kerapatan")

	updt, err := conn.Prepare("update image set kerapatan=? where nama=?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = updt.Exec(kerapatan, nama)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
		})
		fmt.Print(err.Error())
	}
	defer updt.Close()
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "sukses update kerapatan image",
	})
}

func DeleteImage(c *gin.Context) {
	nama := c.Query("nama")
	dlt, err := conn.Prepare("delete from image where nama= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = dlt.Exec(nama)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"message": "delete image ok gan",
	})
}
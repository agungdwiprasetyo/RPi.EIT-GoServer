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

	rows, err := db.Query("select * from image;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&image.Nama, &image.Id_data, &image.Id_algor, &image.Kerapatan, &image.Datetime)
		images = append(images, image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ops, ada kesalahan saat fetch data",
			})
			fmt.Print(err.Error())
			return
		}
	}
	defer rows.Close()
	
	c.JSON(http.StatusOK, images)
}

func PostImage(c *gin.Context) {
	nama := c.PostForm("filename")
	idData := c.PostForm("id_data")
	idAlgor := c.PostForm("id_algor")
	kerapatan := c.PostForm("kerapatan")

	if nama == "" || kerapatan == "" || idData=="" || idAlgor=="" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "tidak ada parameter yang dikirim",
		})
		return
	}

	masuk, err := db.Prepare("insert into image (nama, id_data, id_algor, kerapatan) values(?,?,?,?);")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	_, err = masuk.Exec(nama, idData, idAlgor, kerapatan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat eksekusi query",
		})
		fmt.Print(err.Error())
		return
	}
	defer masuk.Close()

	c.JSON(http.StatusCreated, gin.H{
		"message": "sukses tambah gambar",
	})
}

func PutImage(c *gin.Context) {
	nama := c.PostForm("filename")
	kerapatan := c.PostForm("kerapatan")

	if nama == "" || kerapatan == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "tidak ada parameter yang dikirim",
		})
		return
	}

	updt, err := db.Prepare("update image set kerapatan=? where nama=?;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	_, err = updt.Exec(kerapatan, nama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat eksekusi query",
		})
		fmt.Print(err.Error())
		return
	}
	defer updt.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "sukses update kerapatan image",
	})
}

func DeleteImage(c *gin.Context) {
	nama := c.Param("id")
	if nama == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "tidak ada parameter yang dikirim",
		})
		return
	}

	dlt, err := db.Prepare("delete from image where nama=?;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	_, err = dlt.Exec(nama)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat eksekusi query",
		})
		fmt.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete image sukses gan",
	})
}
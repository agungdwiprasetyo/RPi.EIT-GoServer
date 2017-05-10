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

	rows, err := db.Query("select * from data_ukur;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	for rows.Next() {
		rows.Scan(&data.Id_data, &data.Nama, &data.Filename, &data.Deskripsi, &data.Model, &data.Citra, &data.Arus, &data.Datetime)
		alldata = append(alldata, data)
	}
	defer rows.Close()

	c.JSON(http.StatusOK, alldata)
}

func PostData(c *gin.Context) {
	namaData := c.PostForm("nama_data")
	filename := c.PostForm("filename")
	arusInjeksi := c.PostForm("arus_injeksi")
	deskripsi := c.PostForm("deskripsi")

	if namaData == "" || filename == "" || arusInjeksi=="" || deskripsi=="" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "tidak ada parameter yang dikirim",
		})
		return
	}

	masuk, err := db.Prepare("insert into data_ukur (nama_data, filename, arus_injeksi, deskripsi) values(?,?,?, ?);")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	_, err = masuk.Exec(namaData, filename, arusInjeksi, deskripsi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat eksekusi query",
		})
		fmt.Print(err.Error())
		return
	}
	defer masuk.Close()

	c.JSON(http.StatusCreated, gin.H{
		"message": "sukses tambah data",
	})
}

func PutData(c *gin.Context) {
	var kueri string
	var param string

	citra := c.PostForm("citra")
	idData := c.PostForm("id_data")
	model := c.PostForm("model")
	alamatData := c.PostForm("alamat_data")

	if citra!="" {
		if citra=="apus" {
			citra = ""
		}
		param = citra
		kueri = "update data_ukur set citra=? where id_data=?;"
	}else if model!="" {
		param = model
		kueri = "update data_ukur set model=? where id_data=?;"
	}else {
		param = alamatData
		kueri = "update data_ukur set alamat_data=? where id_data=?;"
	}

	updt, err := db.Prepare(kueri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	_, err = updt.Exec(param, idData)
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

func DeleteData(c *gin.Context) {
	idData := c.Param("id")
	if idData == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "tidak ada parameter yang dikirim",
		})
		return
	}

	dlt, err := db.Prepare("delete from data_ukur where id_data=?;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat query data",
		})
		fmt.Print(err.Error())
		return
	}

	_, err = dlt.Exec(idData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ops, ada kesalahan saat eksekusi query",
		})
		fmt.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete data sukses gan",
	})
}
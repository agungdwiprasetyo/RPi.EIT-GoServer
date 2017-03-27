package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/googollee/go-socket.io"
)

type Image struct {
	Id_image int
	Id_data int
	Nama string
	Id_algor string
	Kerapatan float64
	Datetime string
}

type Data struct {
	Id_data int
	Nama string
	Filename string
	Arus float64
	Datetime string
}

type Algor struct {
	Id_algor string
	Nama string
}

var socket * socketio.Server

func main() {
	// konek database
	db, err := sql.Open("mysql", "root:fakepassword@tcp(127.0.0.1:3306)/eit")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	
	// init socket.io
	// socket, errsock := socketio.NewServer(nil)
	// if errsock != nil {
	// 	fmt.Print(errsock)
	// }

	// routing API
	router := gin.Default()
	router.GET("/image", func(c *gin.Context) {
		var (
			image  Image
			images []Image
		)
		rows, err := db.Query("select * from image;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&image.Id_image, &image.Id_data, &image.Nama, &image.Id_algor, &image.Kerapatan, &image.Datetime)
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
	})

	router.GET("/data", func(c *gin.Context) {
		var (
			data  Data
			alldata []Data
		)
		rows, err := db.Query("select * from data_ukur;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&data.Id_data, &data.Nama, &data.Filename, &data.Arus, &data.Datetime)
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
	})
	router.GET("/algor", func(c *gin.Context) {
		var (
			algor  Algor
			allAlgor []Algor
		)
		rows, err := db.Query("select * from algoritma;")
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
	})

	// router.GET("/socket.io/", socketHandler)
	router.Run(":3000")
}

func socketHandler(c *gin.Context) {
	socket.On("connection", func(so socketio.Socket) {
		fmt.Println("on connection")
		so.Join("RPi.EIT")
		so.On("runReconstruction", func(msg string) {
			fmt.Println("this reconstruction")
		})
		so.On("raspiConnect", func(msg string) {
			fmt.Println("raspi connected")
		})
		so.On("disconnection", func() {
			fmt.Println("on disconnect")
		})
	})
	socket.On("error", func(so socketio.Socket, errsock error) {
		fmt.Println("error:", errsock)
	})
}
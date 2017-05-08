package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"./models"
	"./routers"
	"./database"
)

func main() {
	fmt.Println("RESTful API written in Golang (dipecah dalam bentuk modul)")
	fmt.Println("---- Developed by Agung Dwi Prasetyo ----")
	
	// konek database
	db, err := database.DB_Connect()
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	// init database
	models.SetDB(db)

	// routing API
	router := gin.Default()
	routers.Data(router)
	routers.Image(router)
	routers.Algor(router)

	router.Run(":3456")
	http.Handle("/", http.FileServer(http.Dir("/public")))
}

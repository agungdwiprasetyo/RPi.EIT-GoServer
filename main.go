package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"./routers"
	"./database"
)

func main() {
	fmt.Println("RESTful API written in Golang (dipecah dalam bentuk modul)")
	fmt.Println("---- Developed by Agung Dwi Prasetyo ----")
	
	// konek database
	database.Connect()

	http.Handle("/", http.FileServer(http.Dir("./public")))

	// routing API
	router := gin.Default()
	routers.Data(router)
	routers.Image(router)
	routers.Algor(router)

	router.Run(":3456")
}

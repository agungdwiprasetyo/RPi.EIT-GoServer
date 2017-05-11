package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"./routers"
	"./database"
	"./auth"
)

func main() {
	fmt.Println("RESTful API written in Golang (dipecah dalam bentuk modul)")
	fmt.Println("---- Developed by Agung Dwi Prasetyo ----")
	
	// konek database
	database.Connect()

	// routing API
	router := gin.Default()
	auth.Authenticate(router)
	
	routers.Data(router)
	routers.Image(router)
	routers.Algor(router)

	// router.Run(":3456")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":3456", router)
}

package main

import (
	"fmt"
	// "net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"

	"./routers"
	"./database"
	"./auth"
)

func main() {
	fmt.Println("RESTful API / Back-End side written in Golang (dipecah dalam bentuk modul)")
	fmt.Println("---- Developed by Agung Dwi Prasetyo ----")
	
	// konek database
	database.Connect()

	// init app
	app := gin.Default()

	// otentikasi
	auth.Authenticate(app)

	// init routing API
	routers.InitAPI(app)

	// serve static html Front-End side
	app.Use(static.Serve("/", static.LocalFile("./frontend", true)))

	// run server at port 3456
	app.Run(":3456")
}

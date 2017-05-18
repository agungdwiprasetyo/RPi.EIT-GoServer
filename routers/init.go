package routers

import (
	"fmt"
	"../auth"
	"github.com/gin-gonic/gin"
)

var (
	SecretKey = auth.GetSecretKey()
)

func InitAPI(app *gin.Engine) {
	api := app.Group("/api")
	api.Use(AuthMiddleware())
	{
		Data(api)
		Image(api)
		Algor(api)
	}	
}

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		valid := auth.CheckAuthorization(c)
		if valid {
			fmt.Println("valid")
			c.Next()
		} else {
			c.Abort()
			fmt.Println("gak valid")
		}
	}
}
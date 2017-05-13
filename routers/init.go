package routers

import (
	"../auth"
	"github.com/gin-gonic/gin"
)

var (
	SecretKey = auth.GetSecretKey()
)

func InitAPI(app *gin.Engine) {
	api := app.Group("/api")
	Data(api)
	Image(api)
	Algor(api)
}
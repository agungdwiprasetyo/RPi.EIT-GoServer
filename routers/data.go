package routers

import (
	"github.com/gin-gonic/gin"
	"../models"
)

func Data(api *gin.RouterGroup){
	api.GET("/data", func(c *gin.Context) {
		models.GetData(c)
	})

	api.POST("/data", func(c *gin.Context) {
		models.PostData(c)
	})

	api.PUT("/data", func(c *gin.Context) {
		models.PutData(c)
	})

	api.DELETE("/data/:id", func(c *gin.Context) {
		models.DeleteData(c)
	})
}
package routers

import (
	"github.com/gin-gonic/gin"
	"../models"
)

func Data(router *gin.Engine){
	router.GET("/data", func(c *gin.Context) {
		models.GetData(c)
	})

	router.POST("/data", func(c *gin.Context) {
		models.PostData(c)
	})

	router.PUT("/data", func(c *gin.Context) {
		models.PutData(c)
	})

	router.DELETE("/data", func(c *gin.Context) {
		models.DeleteData(c)
	})
}
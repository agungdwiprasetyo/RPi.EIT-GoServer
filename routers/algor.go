package routers

import (
	"github.com/gin-gonic/gin"
	"../models"
)

func Algor(router *gin.Engine){
	router.GET("/algor", func(c *gin.Context) {
		models.GetAlgor(c)
	})

	router.POST("/algor", func(c *gin.Context) {
		models.PostAlgor(c)
	})

	router.PUT("/algor", func(c *gin.Context) {
		models.PutAlgor(c)
	})

	router.DELETE("/algor", func(c *gin.Context) {
		models.DeleteAlgor(c)
	})
}
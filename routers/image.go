package routers

import (
	"github.com/gin-gonic/gin"
	"../models"
)

func Image(router *gin.Engine){
	router.GET("/image", func(c *gin.Context) {
		models.GetImage(c)
	})

	router.POST("/image", func(c *gin.Context) {
		models.PostImage(c)
	})

	router.PUT("/image", func(c *gin.Context) {
		models.PutImage(c)
	})

	router.DELETE("/image", func(c *gin.Context) {
		models.DeleteImage(c)
	})
}
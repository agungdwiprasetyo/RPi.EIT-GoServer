package routers

import (
	"github.com/gin-gonic/gin"
	"../models"
)

func Image(api *gin.RouterGroup){
	api.GET("/image", func(c *gin.Context) {
		models.GetImage(c)
	})

	api.POST("/image", func(c *gin.Context) {
		models.PostImage(c)
	})

	api.PUT("/image", func(c *gin.Context) {
		models.PutImage(c)
	})

	api.DELETE("/image/:id", func(c *gin.Context) {
		models.DeleteImage(c)
	})
}
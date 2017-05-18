package routers

import (
	"github.com/gin-gonic/gin"
	"../models"
)

func Algor(api *gin.RouterGroup){
	api.GET("/algor", func(c *gin.Context) {
		models.GetAlgor(c)
	})

	api.POST("/algor", func(c *gin.Context) {
		models.PostAlgor(c)
	})

	api.PUT("/algor", func(c *gin.Context) {
		models.PutAlgor(c)
	})

	api.DELETE("/algor/:id", func(c *gin.Context) {
		models.DeleteAlgor(c)
	})
}
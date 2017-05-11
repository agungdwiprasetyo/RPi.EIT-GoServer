package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"../models"
)

func Algor(router *gin.Engine){
	router.GET("/algor", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.GetAlgor(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.POST("/algor", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.PostAlgor(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.PUT("/algor", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.PutAlgor(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.DELETE("/algor/:id", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.DeleteAlgor(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})
}
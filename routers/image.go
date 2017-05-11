package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"../models"
)

func Image(router *gin.Engine){
	router.GET("/image", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.GetImage(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.POST("/image", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.PostImage(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.PUT("/image", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.PutImage(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.DELETE("/image/:id", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.DeleteImage(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})
}
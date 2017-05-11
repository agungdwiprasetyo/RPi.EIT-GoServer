package routers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"../models"
)

func Data(router *gin.Engine){
	router.GET("/data", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.GetData(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.POST("/data", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.PostData(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.PUT("/data", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.PutData(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})

	router.DELETE("/data/:id", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")

		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			models.DeleteData(c)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid",
			})
		}
	})
}
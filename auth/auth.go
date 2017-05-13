package auth

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"

	"../models"
)

const (
	SecretKey = "jancukkabeh"
)

func GetSecretKey() string{
	return SecretKey
}

func Authenticate(app *gin.Engine) {
	app.POST("/login", func(c *gin.Context) {
		user := c.PostForm("username")
		pass := c.PostForm("password")
		var (
			login  models.Login
		)

		login = models.AuthLogin(user)

		if pass==login.Password {
			token := jwt.New(jwt.GetSigningMethod("HS256"))
			claims := token.Claims.(jwt.MapClaims)
			claims["userid"] = login.Username
			claims["tipe"] = login.Tipe
			claims["nama"] = login.Nama

			// Expire in 30 menit
			claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
			tokenString, err := token.SignedString([]byte(SecretKey))
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "login success",
				"token": tokenString,
				"user": claims,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "username or password invalid",
			})
		}
	})

	app.GET("/cek", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")
		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.JSON(http.StatusOK, gin.H{
				"username": claims["userid"],
				"akses": claims["tipe"],
				"nama": claims["nama"],
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid/expired",
			})
		}
	})
}
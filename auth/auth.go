package auth

import (
	"time"
	// "fmt"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"encoding/base64"

	"../models"
)

const (
	SecretKey = "backendserviceforelectricalimpedancetomography"
	expireMinutes = 30
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

		// compare empty struct (no username registered)
		if (models.Login{})==login {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "username tidak terdaftar",
			})
			return
		}

		if pass==login.Password {
			token := jwt.New(jwt.GetSigningMethod("HS256"))
			claims := token.Claims.(jwt.MapClaims)
			claims["userid"] = login.Username
			claims["tipe"] = login.Tipe
			claims["nama"] = login.Nama

			// Expire in 30 menit
			claims["exp"] = time.Now().Add(time.Minute * expireMinutes).Unix()
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
		checkToken := CheckAuthorization(c)
		if checkToken {
			c.JSON(http.StatusOK, gin.H{
				"message" : "mantab, sukses",
			})
		}

	})

	app.GET("/logout", func(c *gin.Context) {
		valid := CheckAuthorization(c)
		if valid {
			c.JSON(http.StatusOK, gin.H{
				"message": "logout success",
			})
		}
	})
}

func CheckAuthorization(c *gin.Context) bool{
	authorization := c.Request.Header.Get("Authorization")
	if(authorization==""){
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "no authorization",
		})
		return false
	}

	tmp1 := strings.Split(authorization," ") // hilangin string Basic

	authDecode, _  := base64.StdEncoding.DecodeString(tmp1[1]) // decode authorization
	basicAuth := string(authDecode[:])

	tmp2 := strings.Split(basicAuth,":")
	name := tmp2[0]
	headToken := tmp2[1]

	token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token invalid",
		})
		return false
	}

	claims := token.Claims.(jwt.MapClaims)
	if token.Valid && name==claims["userid"] {
		return true
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token expired",
		})
		return false
	}
}

func GetRemainingToken(waktu interface{}) int{
	if validity, ok := waktu.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds())
		}
	}
	return expireMinutes
}
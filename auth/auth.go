package auth

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"

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
		headToken := c.Request.Header.Get("token")
		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		// claims := token.Claims.(jwt.MapClaims)
		if err == nil && token.Valid {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid/expired",
				"token": token,
			})
		}
	})

	app.GET("/logout", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")
		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			rem := GetRemainingToken(claims["exp"])
			c.JSON(http.StatusOK, gin.H{
				"message": "logout sukses",
				"exp": rem,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "token invalid/expired",
			})
		}
	})
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
package auth

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

const (
	ValidUser = "Agung"
	ValidPass = "fake"
	SecretKey = "jancukkabeh"
)

func GetSecretKey() string{
	return SecretKey
}

func Authenticate(router *gin.Engine) {
	router.POST("/login", func(c *gin.Context) {
		user := c.PostForm("username")
		pass := c.PostForm("password")
		if user==ValidUser && pass==ValidPass {
			token := jwt.New(jwt.GetSigningMethod("HS256"))
			claims := token.Claims.(jwt.MapClaims)
			claims["userid"] = user
			claims["tipe"] = "admin"

			// Expire in 30 menit
			claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
			tokenString, err := token.SignedString([]byte(SecretKey))
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"token": tokenString,	
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "username or password invalid",
			})
		}
	})

	router.GET("/cek", func(c *gin.Context) {
		headToken := c.Request.Header.Get("token")
		token, err := jwt.Parse(headToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.JSON(http.StatusOK, gin.H{
				"userid": claims["userid"],
				"akses": claims["tipe"],
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "token invalid",
			})
		}
	})
}
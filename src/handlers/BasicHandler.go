package handlers

import (
	"errors"
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/go_gin_govendor/src/models"
)

var petterToken string = "secret*#key#*for*#AES&encryption"

func ApikeyValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.Request.Header.Get("Authorization")
		apikey := c.Request.Header.Get("apikey")

		// log.Println(apikey)
		if apikey == "" {
			ResponseError(c, http.StatusForbidden, "Apikey required")
			c.Error(errors.New("No auth Apikey"))
			c.Abort()
			return
		}
		c.Next()
	}
}
func TokenValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.Request.Header.Get("Authorization")
		auth := c.Request.Header.Get("Authorization")

		// log.Println(apikey)
		if auth == "" {
			ResponseError(c, http.StatusForbidden, "Authorization required")
			c.Error(errors.New("No auth Authorization"))
			c.Abort()
			return
		}

		if !ValidatorJWT(c.Request.Header.Get("Authorization"), c) {
			c.Error(errors.New("No auth token sent"))
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"msg":    "Auth token is invalid",
			})
			c.Abort()
			return
		}

		c.Set("token", c.Request.Header.Get("Authorization"))
		c.Next()
	}
}
func GenerateJWT(userid int) (string, error) {
	myKey := []byte("secret*#key#*for*#AES&encryption")
	claims := models.JwtStruct{
		userid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "CRUD",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(myKey)
	if err != nil {
		return "", nil
	}

	return signedString, nil
}

func ValidatorJWT(token string, c *gin.Context) bool {
	newToken, _ := jwt.ParseWithClaims(token, &models.JwtStruct{}, func(newToken *jwt.Token) (interface{}, error) {
		return []byte(petterToken), nil
	})
	if claims, ok := newToken.Claims.(*models.JwtStruct); ok && newToken.Valid {
		c.Set("iduser", claims.UserId)
		return true
	}
	return false
}

func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"status": "false",
		"msg":    msg,
	})
}

func ResponseSuccess(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}

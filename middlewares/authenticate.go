package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *gin.Context) {

	var responseCode int = 401

	tokenString := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)

	if tokenString == "" {

		c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseOnlyMessage{
			Code:    responseCode,
			Message: "Request tidak valid",
		})

	} else {

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			var client models.User

			configs.DB.Raw("SELECT a.kar_id, a.kar_nama, a.kar_pass, b.kp_ket FROM tm_karyawan a JOIN tr_kota_pelayanan b ON a.kp_kode = b.kp_kode WHERE a.kar_id = ? AND a.grup_id = '020'", claims["sub"]).Find(&client)

			if client.KarId == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseOnlyMessage{
					Code:    responseCode,
					Message: "Token tidak terdaftar pada Pengguna",
				})
			}

			c.Next()

		} else if errors.Is(err, jwt.ErrTokenExpired) {

			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseOnlyMessage{
				Code:    responseCode,
				Message: "Token kedaluwarsa",
			})

		} else {

			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ResponseOnlyMessage{
				Code:    responseCode,
				Message: "Token tidak valid",
			})

		}

	}

}

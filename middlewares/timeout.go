package middlewares

import (
	"net/http"
	"time"

	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Timeout() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(65*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(TimeoutResponse),
	)
}

func TimeoutResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, models.ResponseOnlyMessage{
		Code:    408,
		Message: "Request timeout",
	})
}

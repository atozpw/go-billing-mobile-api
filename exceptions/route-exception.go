package exceptions

import (
	"net/http"

	"github.com/atozpw/go-billing-mobile-api/models"
	"github.com/gin-gonic/gin"
)

func RouteException(c *gin.Context) {

	c.JSON(http.StatusNotFound, models.ResponseOnlyMessage{
		Code:    404,
		Message: "Resource tidak ditemukan",
	})

}

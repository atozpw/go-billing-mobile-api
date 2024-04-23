package configs

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Mode() {

	if os.Getenv("APP_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

}

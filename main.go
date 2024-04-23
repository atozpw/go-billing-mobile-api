package main

import (
	"os"

	"github.com/atozpw/go-billing-mobile-api/configs"
	"github.com/atozpw/go-billing-mobile-api/helpers"
	"github.com/atozpw/go-billing-mobile-api/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	configs.LoadEnvironment()
	configs.ConnectDatabase()
	configs.Logging()
	configs.Mode()

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(helpers.LoggerFormatter))
	router.Use(gin.Recovery())
	router.Use(middlewares.Cors())
	router.Use(middlewares.Timeout())
	Routes(router)
	router.Run(os.Getenv("APP_URL"))

}

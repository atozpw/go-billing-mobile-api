package main

import (
	"github.com/atozpw/go-billing-mobile-api/controllers"
	"github.com/atozpw/go-billing-mobile-api/exceptions"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.POST("/login", controllers.Login)
		v1.POST("/register", controllers.Register)
		v1.GET("/customers/:id", controllers.CustomerFind)
		v1.GET("/customers/:id/bills", controllers.CustomerBills)
		v1.GET("/payments", controllers.PaymentIndex)
		v1.POST("/payments", controllers.PaymentStore)
		v1.GET("/profile", controllers.Profile)
		v1.POST("/change-password", controllers.ChangePassword)
		v1.Static("/storage", "./storages/public")
	}

	router.NoRoute(exceptions.RouteException)

}

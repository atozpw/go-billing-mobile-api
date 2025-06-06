package main

import (
	"net/http"

	"github.com/atozpw/go-billing-mobile-api/controllers"
	"github.com/atozpw/go-billing-mobile-api/exceptions"
	"github.com/atozpw/go-billing-mobile-api/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.POST("/login", controllers.Login)
		v1.POST("/register", middlewares.Auth, controllers.Register)
		v1.GET("/customers", middlewares.Auth, controllers.CustomerIndex)
		v1.GET("/customers/:id", middlewares.Auth, controllers.CustomerFind)
		v1.GET("/customers/:id/bills", middlewares.Auth, controllers.CustomerBills)
		v1.GET("/bills", middlewares.Auth, controllers.BillIndex)
		v1.GET("/bills/:id", middlewares.Auth, controllers.BillFind)
		v1.GET("/payments", middlewares.Auth, controllers.PaymentIndex)
		v1.GET("/payments/:id", middlewares.Auth, controllers.PaymentFind)
		v1.POST("/payments", middlewares.Auth, controllers.PaymentStore)
		v1.GET("/reports", middlewares.Auth, controllers.ReportIndex)
		v1.GET("/profile", middlewares.Auth, controllers.Profile)
		v1.POST("/change-password", middlewares.Auth, controllers.ChangePassword)
		v1.POST("/receipt/whatsapp", middlewares.Auth, controllers.ReceiptToWhatsapp)
		v1.StaticFS("/storage", http.Dir("./storages/public"))
	}

	router.NoRoute(exceptions.RouteException)

}

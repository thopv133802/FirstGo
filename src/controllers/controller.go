package controllers

import (
	"GoFirst/src/domain/services"
	"github.com/gin-gonic/gin"
	"log"
)

type CommonController interface {
	RunAPI(address string) error
}

type CommonControllerImpl struct {
	Service services.CommonServiceImpl
}

func (controller *CommonControllerImpl) RunAPI(address string) error {
	routes := gin.Default()
	service := controller.Service
	routes.GET("/products", service.GetProducts)
	routes.GET("/promos", service.GetPromos)
	userGroup := routes.Group("/user")
	{
		userGroup.POST("/:id/signout", service.SignOut)
		userGroup.GET("/:id/orders", service.GetOrders)
	}
	usersGroup := routes.Group("/users")
	{
		usersGroup.POST("/signin", service.SignIn)
		usersGroup.POST("", service.AddUser)
		usersGroup.POST("/charge", service.Charge)
	}
	err := routes.Run(address)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
package routes

import (
	"explore-aks-backend-app-air/src/controllers"
	"explore-aks-backend-app-air/src/services"

	"github.com/gin-gonic/gin"
)

func setupCustomerRoutes(router *gin.RouterGroup) {
	controller := controllers.CustomerController{
		ServiceCfg: services.CustomerService.CRUDServiceConfig,
	}

	router.GET("/customers/:id", controller.GetCustomer)
	router.GET("/customers", controller.GetListCustomers)
	router.POST("/customers", controller.CreateCustomer)
	router.PATCH("/customers/:id", controller.UpdateCustomer)
	router.DELETE("/customers/:id", controller.DeleteCustomer)
}

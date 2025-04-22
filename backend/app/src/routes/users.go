package routes

import (
	"explore-aks-backend-app-air/src/controllers"
	"explore-aks-backend-app-air/src/services"

	"github.com/gin-gonic/gin"
)

func setupUserRoutes(router *gin.RouterGroup) {
	controller := controllers.UserController{
		ServiceCfg: services.UserService.CRUDServiceConfig,
	}

	router.GET("/users/:id", controller.GetUser)
	router.GET("/users", controller.GetListUsers)
	router.POST("/users", controller.CreateUser)
	router.PATCH("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}

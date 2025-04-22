package routes

import (
	"explore-aks-backend-app-air/src/controllers"
	"explore-aks-backend-app-air/src/services"

	"github.com/gin-gonic/gin"
)

func setupProjecttimeRoutes(router *gin.RouterGroup) {
	controller := controllers.ProjecttimeController{
		ServiceCfg: services.ProjecttimeService.CRUDServiceConfig,
	}

	router.GET("/projecttimes/:id", controller.GetProjecttime)
	router.GET("/projecttimes", controller.GetListProjecttimes)
	router.POST("/projecttimes", controller.CreateProjecttime)
	router.PATCH("/projecttimes/:id", controller.UpdateProjecttime)
	router.DELETE("/projecttimes/:id", controller.DeleteProjecttime)
}

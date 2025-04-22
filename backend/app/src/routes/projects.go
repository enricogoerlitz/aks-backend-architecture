package routes

import (
	"explore-aks-backend-app-air/src/controllers"
	"explore-aks-backend-app-air/src/services"

	"github.com/gin-gonic/gin"
)

func setupProjectRoutes(router *gin.RouterGroup) {
	controller := controllers.ProjectController{
		ServiceCfg: services.ProjectService.CRUDServiceConfig,
	}

	router.GET("/projects/:id", controller.GetProject)
	router.GET("/projects", controller.GetListProjects)
	router.POST("/projects", controller.CreateProject)
	router.PATCH("/projects/:id", controller.UpdateProject)
	router.DELETE("/projects/:id", controller.DeleteProject)
}

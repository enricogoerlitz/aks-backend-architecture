package controllers

import (
	"explore-aks-backend-app-air/src/libs/crud"
	"explore-aks-backend-app-air/src/models"
	"explore-aks-backend-app-air/src/services"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	Service services.ProjectCRUDService
}

func (pc *ProjectController) GetProject(c *gin.Context) {
	id := c.Param("id")
	crud.HandleGet[models.ProjectResponse](c, id, nil)
}

func (pc *ProjectController) GetListProjects(c *gin.Context) {
	crud.HandleGetList[models.ProjectResponse](c, pc.Service.Config, nil)
}

func (pc *ProjectController) CreateProject(c *gin.Context) {
	crud.HandlePost[models.ProjectPOST, models.ProjectResponse](c, pc.Service.Config)
}

func (pc *ProjectController) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	crud.HandlePatch[models.ProjectModel, models.ProjectPATCH](c, pc.Service.Config, id)
}

func (pc *ProjectController) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	crud.HandleDelete[models.ProjectModel](c, id)
}

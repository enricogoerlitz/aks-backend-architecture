package controllers

import (
	"explore-aks-backend-app-air/src/libs/crud"
	"explore-aks-backend-app-air/src/models"
	"explore-aks-backend-app-air/src/services"

	"github.com/gin-gonic/gin"
)

type ProjecttimeController struct {
	Service services.ProjecttimeCRUDService
}

func (ptc *ProjecttimeController) GetProjecttime(c *gin.Context) {
	id := c.Param("id")
	ptc.Service.HandleGet(c, id)
}

func (ptc *ProjecttimeController) GetListProjecttimes(c *gin.Context) {
	ptc.Service.HandleGetList(c)
}

func (ptc *ProjecttimeController) CreateProjecttime(c *gin.Context) {
	crud.HandlePost[models.ProjecttimePOST, models.ProjecttimeResponse](c, ptc.Service.Config)
}

func (ptc *ProjecttimeController) UpdateProjecttime(c *gin.Context) {
	id := c.Param("id")
	crud.HandlePatch[models.ProjecttimeModel, models.ProjecttimePATCH](c, ptc.Service.Config, id)
}

func (ptc *ProjecttimeController) DeleteProjecttime(c *gin.Context) {
	id := c.Param("id")
	crud.HandleDelete[models.ProjecttimeModel](c, id)
}

package controllers

import (
	"explore-aks-backend-app-air/src/libs/crud"
	"explore-aks-backend-app-air/src/models"

	"github.com/gin-gonic/gin"
)

type ProjecttimeController struct {
	ServiceCfg *crud.CRUDServiceConfig
}

func (ptc *ProjecttimeController) GetProjecttime(c *gin.Context) {
	id := c.Param("id")
	crud.HandleGet[models.ProjecttimeResponse](c, id, nil)
}

func (ptc *ProjecttimeController) GetListProjecttimes(c *gin.Context) {
	crud.HandleGetList[models.ProjecttimeResponse](c, ptc.ServiceCfg, nil)
}

func (ptc *ProjecttimeController) CreateProjecttime(c *gin.Context) {
	crud.HandlePost[models.ProjecttimePOST, models.ProjecttimeResponse](c, ptc.ServiceCfg)
}

func (ptc *ProjecttimeController) UpdateProjecttime(c *gin.Context) {
	id := c.Param("id")
	crud.HandlePatch[models.ProjecttimeModel, models.ProjecttimePATCH](c, ptc.ServiceCfg, id)
}

func (ptc *ProjecttimeController) DeleteProjecttime(c *gin.Context) {
	id := c.Param("id")
	crud.HandleDelete[models.ProjecttimeModel](c, id)
}

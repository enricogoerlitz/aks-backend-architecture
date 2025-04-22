package services

import (
	"explore-aks-backend-app-air/src/database"
	"explore-aks-backend-app-air/src/libs/crud"
	"explore-aks-backend-app-air/src/models"

	"github.com/gin-gonic/gin"
)

type ProjecttimeCRUDService struct {
	*crud.CRUDServiceConfig
}

func (s *ProjecttimeCRUDService) HandleGet(c *gin.Context, id interface{}) {
	isDetailRoute := c.Query("detail") == "true"

	if isDetailRoute {
		dbQuery := database.DBRead.Preload("User").Preload("Project")
		crud.HandleGet[models.ProjecttimeResponseDetail](c, id, dbQuery)
	} else {
		crud.HandleGet[models.ProjecttimeResponse](c, id, nil)
	}
}

func (s *ProjecttimeCRUDService) HandleGetList(c *gin.Context) {
	isDetailRoute := c.Query("detail") == "true"

	if isDetailRoute {
		dbQuery := database.DBRead.Preload("User").Preload("Project")
		crud.HandleGetList[models.ProjecttimeResponseDetail](c, s.CRUDServiceConfig, dbQuery)
	} else {
		crud.HandleGetList[models.ProjecttimeResponse](c, s.CRUDServiceConfig, nil)
	}
}

var userForeignKeyColumn = crud.ForeignKeyColumn{
	Model: &models.UserModel{},
	Field: "UserID",
}

var projectForeignKeyColumn = crud.ForeignKeyColumn{
	Model: &models.ProjectModel{},
	Field: "ProjectID",
}

var ProjecttimeService = ProjecttimeCRUDService{
	CRUDServiceConfig: &crud.CRUDServiceConfig{
		UniqueColumns:         []string{},
		UniqueColumnsTogether: [][]string{},
		ForeignKeyColumns: []crud.ForeignKeyColumn{
			userForeignKeyColumn,
			projectForeignKeyColumn,
		},
		ReadOnlyFields:        []string{},
		PaginationMaxPageSize: 100,
	},
}

package services

import (
	"explore-aks-backend-app-air/src/libs/crud"
	"explore-aks-backend-app-air/src/models"
)

type ProjectCRUDService struct {
	*crud.CRUDServiceConfig
}

var customerForeignKeyColumn = crud.ForeignKeyColumn{
	Model: &models.CustomerModel{},
	Field: "CustomerID",
}

var ProjectService = ProjectCRUDService{
	CRUDServiceConfig: &crud.CRUDServiceConfig{
		UniqueColumns:         []string{"Name"},
		UniqueColumnsTogether: [][]string{},
		ForeignKeyColumns: []crud.ForeignKeyColumn{
			customerForeignKeyColumn,
		},
		ReadOnlyFields:        []string{},
		PaginationMaxPageSize: 100,
	},
}

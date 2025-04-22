package services

import "explore-aks-backend-app-air/src/libs/crud"

type CustomerCRUDService struct {
	Config *crud.CRUDServiceConfig
}

var CustomerService = CustomerCRUDService{
	Config: &crud.CRUDServiceConfig{
		UniqueColumns:         []string{"Name"},
		UniqueColumnsTogether: [][]string{},
		ForeignKeyColumns:     []crud.ForeignKeyColumn{},
		ReadOnlyFields:        []string{},
		PaginationMaxPageSize: 100,
	},
}

package services

import "explore-aks-backend-app-air/src/libs/crud"

type CustomerCRUDService struct {
	*crud.CRUDServiceConfig
}

var CustomerService = CustomerCRUDService{
	CRUDServiceConfig: &crud.CRUDServiceConfig{
		UniqueColumns:         []string{"Name"},
		UniqueColumnsTogether: [][]string{},
		ForeignKeyColumns:     []crud.ForeignKeyColumn{},
		ReadOnlyFields:        []string{},
		PaginationMaxPageSize: 100,
	},
}

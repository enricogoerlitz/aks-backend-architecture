package services

import "explore-aks-backend-app-air/src/libs/crud"

type UserCRUDService struct {
	*crud.CRUDServiceConfig
}

var UserService = UserCRUDService{
	CRUDServiceConfig: &crud.CRUDServiceConfig{
		UniqueColumns: []string{},
		UniqueColumnsTogether: [][]string{
			{"Firstname", "Lastname"},
		},
		ForeignKeyColumns:     []crud.ForeignKeyColumn{},
		ReadOnlyFields:        []string{},
		PaginationMaxPageSize: 100,
	},
}

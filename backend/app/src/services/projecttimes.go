package services

import (
	"explore-aks-backend-app-air/src/cache"
	"explore-aks-backend-app-air/src/constants"
	"explore-aks-backend-app-air/src/database"
	"explore-aks-backend-app-air/src/libs/crud"
	"explore-aks-backend-app-air/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

type ProjecttimeCRUDService struct {
	Config *crud.CRUDServiceConfig
}

func (s *ProjecttimeCRUDService) HandleGet(c *gin.Context, id interface{}) {
	isDetailRoute := c.Query("detail") == "true"

	cacheKeyStr := cache.GenerateCacheKey(
		cache.CacheKeyAttribute{
			Attr:  "url",
			Value: c.Request.URL.String(),
		},
		cache.CacheKeyAttribute{
			Attr:  "user_id",
			Value: "1",
		},
	)
	cacheCfg := crud.CachingConfig{
		CacheDB:    cache.Cache,
		Key:        cacheKeyStr,
		Expiration: time.Hour * 1,
		UseCache:   crud.IsHeaderUseCache(c, true),
	}

	if isDetailRoute {
		dbQuery := database.DBRead.Preload("User").Preload("Project")
		crud.HandleGetCached[models.ProjecttimeResponseDetail](c, id, dbQuery, cacheCfg)
	} else {
		crud.HandleGetCached[models.ProjecttimeResponse](c, id, nil, cacheCfg)
	}
}

func (s *ProjecttimeCRUDService) HandleGetList(c *gin.Context) {
	isDetailRoute := c.Query("detail") == "true"

	cacheKeyStr := cache.GenerateCacheKey(
		cache.CacheKeyAttribute{
			Attr:  "url",
			Value: c.Request.URL.String(),
		},
		cache.CacheKeyAttribute{
			Attr:  "user_id",
			Value: "1",
		},
	)
	cacheCfg := crud.CachingConfig{
		CacheDB:    cache.Cache,
		Key:        cacheKeyStr,
		Expiration: time.Duration(constants.ENV_CACHE_TTL_IN_SECONDS) * time.Second,
		UseCache:   crud.IsHeaderUseCache(c, true),
	}

	if isDetailRoute {
		dbQuery := database.DBRead.Preload("User").Preload("Project")
		crud.HandleGetListCached[models.ProjecttimeResponseDetail](c, s.Config, dbQuery, cacheCfg)
	} else {
		crud.HandleGetListCached[models.ProjecttimeResponse](c, s.Config, nil, cacheCfg)
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
	Config: &crud.CRUDServiceConfig{
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

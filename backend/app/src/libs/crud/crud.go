package crud

import (
	"explore-aks-backend-app-air/src/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/**TODOs:
* Jede HTTP-Methode nacheinander durchgehen

- DB-Model => Constrains einbauen/überarbeiten (Unique, NotNull, etc.)


- add model validation

- add queryparam filters
- add additional features
- add sorting

*/

/**
type ProjectTaskID struct {
	ProjectID int
	TaskID    int
}

id could be: id := ProjectTaskID{ProjectID: 1, TaskID: 101}
*/

type ForeignKeyColumn struct {
	Model interface{}
	Field string
}

type CRUDServiceConfig struct {
	UniqueColumns         []string
	UniqueColumnsTogether [][]string
	ForeignKeyColumns     []ForeignKeyColumn
	ReadOnlyFields        []string
	PaginationMaxPageSize int
}

func HandleGet[ResponsePayload any](c *gin.Context, id interface{}, dbQuery *gorm.DB) error {
	var obj ResponsePayload
	dbQuery = GetDBQuery(dbQuery)

	if err := dbQuery.First(&obj, id).Error; err != nil {
		return HandleDatabaseNotFound(c, err)
	}

	c.JSON(http.StatusOK, obj)
	return nil
}

func HandleGetList[ResponsePayload any](c *gin.Context, sc *CRUDServiceConfig, dbQuery *gorm.DB) error {
	var obj []ResponsePayload
	dbQuery = PaginateDBQuery(c, GetDBQuery(dbQuery), sc.PaginationMaxPageSize)

	if err := dbQuery.Find(&obj).Error; err != nil {
		return HandleDatabaseNotFound(c, err)
	}

	c.JSON(http.StatusOK, obj)
	return nil
}

func HandlePostCustomPayload[RequestPayload any, ResponsePayload any](c *gin.Context, sc *CRUDServiceConfig, payload RequestPayload) error {
	payloadValue := ReflectParser.GetReflectValueObject(payload)

	if err := CheckForeignKeysExisting(payloadValue, sc.ForeignKeyColumns, true); err != nil {
		return HandleBadRequest(c, err)
	} else if err := CheckUniqueColumns[RequestPayload](payloadValue, sc.UniqueColumns); err != nil {
		return HandleBadRequest(c, err)
	} else if err := CheckUniqueColumnsTogether[RequestPayload](payloadValue, sc.UniqueColumnsTogether); err != nil {
		return HandleBadRequest(c, err)
	} else if err := CheckUniquePrimaryKey(payloadValue); err != nil {
		return HandleBadRequest(c, err)
	}

	var obj ResponsePayload

	MapStruct(&payload, &obj)
	if err := database.DBWrite.Create(&obj).Error; err != nil {
		return HandleInternalServerError(c, err)
	}

	c.JSON(http.StatusCreated, obj)
	return nil
}

func HandlePost[RequestPayload any, ResponsePayload any](c *gin.Context, sc *CRUDServiceConfig) error {
	var payload RequestPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		return HandlePayloadError(c, err)
	}

	return HandlePostCustomPayload[RequestPayload, ResponsePayload](c, sc, payload)
}

func HandlePatchCustomPayload[BaseModel any, RequestPayload any](c *gin.Context, sc *CRUDServiceConfig, id interface{}, payload RequestPayload) error {
	var obj BaseModel
	if err := database.DBRead.First(&obj, id).Error; err != nil {
		return HandleDatabaseNotFound(c, err)
	}

	payloadValue := ReflectParser.GetReflectValueObject(payload)

	if err := CheckForeignKeysExisting(payloadValue, sc.ForeignKeyColumns, false); err != nil {
		return HandleBadRequest(c, err)
	} else if err := CheckUniqueColumns[RequestPayload](payloadValue, sc.UniqueColumns); err != nil {
		return HandleBadRequest(c, err)
	} else if err := CheckUniqueColumnsTogether[RequestPayload](payloadValue, sc.UniqueColumnsTogether); err != nil {
		return HandleBadRequest(c, err)
	}

	if err := database.DBWrite.Model(&obj).Updates(payload).Error; err != nil {
		return HandleInternalServerError(c, err)
	}

	c.JSON(http.StatusOK, obj)
	return nil
}

func HandlePatch[BaseModel any, RequestPayload any](c *gin.Context, sc *CRUDServiceConfig, id interface{}) error {
	var payload RequestPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		return HandlePayloadError(c, err)
	}

	return HandlePatchCustomPayload[BaseModel, RequestPayload](c, sc, id, payload)
}

func HandleDelete[BaseModel any](c *gin.Context, id ...interface{}) error {
	var obj BaseModel
	if err := database.DBWrite.Delete(obj, id).Error; err != nil {
		return HandleInternalServerError(c, err)
	}

	c.JSON(http.StatusNoContent, nil)
	return nil
}

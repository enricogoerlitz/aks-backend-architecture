package controllers

import (
	"explore-aks-backend-app-air/src/libs/crud"
	"explore-aks-backend-app-air/src/models"
	"explore-aks-backend-app-air/src/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service services.UserCRUDService
}

func (uc *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	crud.HandleGet[models.UserResponse](c, id, nil)
}

func (uc *UserController) GetListUsers(c *gin.Context) {
	crud.HandleGetList[models.UserResponse](c, uc.Service.Config, nil)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// handle different because of the password. but not important for this project!
	crud.HandlePost[models.UserPOST, models.UserResponse](c, uc.Service.Config)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	crud.HandlePatch[models.UserModel, models.UserPATCH](c, uc.Service.Config, id)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	crud.HandleDelete[models.UserModel](c, id)
}

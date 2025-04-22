package controllers

import (
	"explore-aks-backend-app-air/src/libs/crud"
	"explore-aks-backend-app-air/src/models"
	"explore-aks-backend-app-air/src/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	Service services.CustomerCRUDService
}

func (cc *CustomerController) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	crud.HandleGet[models.CustomerResponse](c, id, nil)
}

func (cc *CustomerController) GetListCustomers(c *gin.Context) {
	crud.HandleGetList[models.CustomerResponse](c, cc.Service.Config, nil)
}

func (cc *CustomerController) CreateCustomer(c *gin.Context) {
	crud.HandlePost[models.CustomerPOST, models.CustomerResponse](c, cc.Service.Config)
}

func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	crud.HandlePatch[models.CustomerModel, models.CustomerPATCH](c, cc.Service.Config, id)
}

func (cc *CustomerController) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	crud.HandleDelete[models.CustomerModel](c, id)
}

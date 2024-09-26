package controller

import (
	"github.com/3boku/backend1/service"
	"github.com/3boku/backend1/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShopController struct {
	Service *service.Shop
}

func (controller *ShopController) CreateShop(c *gin.Context) {
	var shop *types.ShopDAO
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, data, err := controller.Service.CreateShop(shop)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": shop, "error": err})
		return
	}
	c.JSON(statusCode, gin.H{"data": data, "error": err})
}

func (controller *ShopController) GetShopByID(c *gin.Context) {
	id := c.Param("id")

	statusCode, shop, err := controller.Service.GetShopByID(id)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": shop, "error": err})
		return
	}
	c.JSON(statusCode, gin.H{"data": shop, "error": err})
}

func (controller *ShopController) UpdateShop(c *gin.Context) {
	id := c.Param("id")
	var shop *types.ShopDAO
	if err := c.ShouldBindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, data, err := controller.Service.UpdateShop(id, shop)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": shop, "error": err})
		return
	}
	c.JSON(statusCode, gin.H{"data": data, "error": err})
}

func (controller *ShopController) DeleteShop(c *gin.Context) {
	id := c.Param("id")
	statusCode, data, err := controller.Service.DeleteShop(id)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": data, "error": err})
		return
	}

	c.JSON(statusCode, gin.H{"data": data, "error": err})
}

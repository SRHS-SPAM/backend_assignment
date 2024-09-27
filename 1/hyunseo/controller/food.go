package controller

import (
	"github.com/3boku/backend1/service"
	"github.com/3boku/backend1/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FoodController struct {
	FoodService *service.FoodService
}

func (f *FoodController) CreateFood(c *gin.Context) {
	shopID := c.Param("shopId")
	var food *types.FoodInput
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, data, err := f.FoodService.CreateFood(shopID, food)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": data, "error": err})
		return
	}
	c.JSON(statusCode, gin.H{"data": data, "error": err})
}

func (f *FoodController) GetFoodByID(c *gin.Context) {
	id := c.Param("id")

	statusCode, shop, err := f.FoodService.SeleteByID(id)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": shop, "error": err})
		return
	}
	c.JSON(statusCode, gin.H{"data": shop, "error": err})
}

func (f *FoodController) UpdateFood(c *gin.Context) {
	id := c.Param("id")
	var food *types.FoodInput
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	statusCode, data, err := f.FoodService.UpdateFood(id, food)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": food, "error": err})
		return
	}
	c.JSON(statusCode, gin.H{"data": data, "error": err})
}

func (f *FoodController) DeleteFood(c *gin.Context) {
	id := c.Param("id")
	statusCode, data, err := f.FoodService.DeleteFood(id)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": data, "error": err})
		return
	}

	c.JSON(statusCode, gin.H{"data": data, "error": err})
}

func (f *FoodController) GetFood(c *gin.Context) {
	shopId := c.Param("shopId")
	statusCode, shops, err := f.FoodService.SeleteAll(shopId)
	if err != nil {
		c.JSON(statusCode, gin.H{"data": shops, "error": err})
		return
	}

	c.JSON(statusCode, gin.H{"data": shops, "error": err})
}

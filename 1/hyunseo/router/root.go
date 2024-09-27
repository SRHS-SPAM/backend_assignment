package router

import (
	"github.com/3boku/backend1/controller"
	"github.com/3boku/backend1/repositories"
	"github.com/3boku/backend1/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	Engine         *gin.Engine
	ShopController *controller.ShopController
	FoodController *controller.FoodController
}

func NewRouter(db *gorm.DB) *Router {
	shop := repositories.ShopRepository{DB: db}
	shopService := service.ShopService{ShopRepository: &shop}
	shopController := controller.ShopController{ShopService: &shopService}
	food := repositories.FoodRepository{DB: db}
	foodService := service.FoodService{FoodRepository: &food}
	foodController := controller.FoodController{FoodService: &foodService}
	r := &Router{
		Engine:         gin.Default(),
		ShopController: &shopController,
		FoodController: &foodController,
	}

	return r
}

func (r *Router) ServerStart() error {
	return r.Engine.Run(":8080")
}

func (r *Router) SetupRoutes() {
	shop := r.Engine.Group("shop")
	{
		shop.PUT("/:id", r.ShopController.UpdateShop)
		shop.DELETE("/:id", r.ShopController.DeleteShop)
		shop.GET("", r.ShopController.GetShop)
		shop.GET("/:id", r.ShopController.GetShopByID)
		shop.POST("", r.ShopController.CreateShop)
	}

	food := r.Engine.Group("food")
	{
		food.PUT("/:shopId/:id", r.FoodController.UpdateFood)
		food.DELETE("/:shopId/:id", r.FoodController.DeleteFood)
		food.GET("/:shopId", r.FoodController.GetFood)
		food.GET("/:shopId/:id", r.FoodController.GetFoodByID)
		food.POST("/:shopId", r.FoodController.CreateFood)
	}
}

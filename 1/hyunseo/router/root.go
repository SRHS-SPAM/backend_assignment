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
}

func NewRouter(db *gorm.DB) *Router {
	shop := repositories.ShopRepository{DB: db}
	shopService := service.ShopService{ShopRepository: &shop}
	shopController := controller.ShopController{ShopService: &shopService}
	r := &Router{
		Engine:         gin.Default(),
		ShopController: &shopController,
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
}

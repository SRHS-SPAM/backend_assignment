package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter() *Router {
	r := &Router{
		engine: gin.Default(),
	}

	return r
}

func (r *Router) ServerStart() error {
	return r.engine.Run(":8080")
}

func (router *Router) SetupRoutes() error {
	router.engine.GET("/shop")
}

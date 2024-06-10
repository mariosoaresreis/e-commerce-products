package api

import (
	"github.com/gin-gonic/gin"
	"products/internal/config"
	"products/internal/services"
)

type ProductHandler struct {
	service services.ProductService
	db      interfaces.DB
}

func (api *API) initRoutes(router *gin.Engine) {
	base := router.Group(config.APIBasePath)

	base.GET("/ping", api.HandlePing)
	base.GET("/livez", func(ctx *gin.Context) {})
	base.GET("/readyz", func(ctx *gin.Context) {})

	v1 := base.Group("v1")
	v1.POST("/products", api.productHandler.Save)
}

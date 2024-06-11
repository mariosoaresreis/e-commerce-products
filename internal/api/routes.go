package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"products/internal/config"
	"products/internal/logger"
)

func validateEnvironmentVariables() {
	envProps := []string{
		"ADDRESS",
		"PORT",
		"USER",
		"PASS",
		"URL",
		"DB_PORT",
		"DB",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(
				fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func (api *API) InitRoutes(router *gin.Engine) {
	base := router.Group(config.APIBasePath)

	base.GET("/ping", api.HandlePing)
	base.GET("/livez", func(ctx *gin.Context) {})
	base.GET("/readyz", func(ctx *gin.Context) {})

	v1 := base.Group("v1")
	v1.POST("/products", api.productHandler.Save)
}

package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"products/internal/logger"
	"products/internal/services"
)

func checkEnvironmentVariables() {
	envProps := []string{
		"ADDRESS",
		"PORT",
		"AUTH_URL",
	}

	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Exiting application...", k))
		}
	}
}
func Start() {
	checkEnvironmentVariables()
	router := *gin.Default()
	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
	handlers := ProductHandlers{services.NewProductService()}
	group := router.Group("api/v1")
	group.POST("/save", handlers.save)
	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	log.Fatalf("%s:%s:%s", address, port, router.Run(port))
}

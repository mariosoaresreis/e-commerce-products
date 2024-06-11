package main

import (
	"github.com/gin-gonic/gin"
	"products/internal/api"
)

func main() {
	router := gin.Default()
	a := api.API{}
	a.InitRoutes(router)
}

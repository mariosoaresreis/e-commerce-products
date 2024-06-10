package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type API struct {
	//config         *config.Config
	swaggerPath    string
	host           string
	port           string
	router         *gin.Engine
	server         *http.Server
	productHandler *ProductHandler
	//logger         logger
}

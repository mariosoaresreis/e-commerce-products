package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"products/internal/services"
)

type ProductHandler struct {
	service services.ProductService
}

func (ph *ProductHandler) Save(ctx context.Context, c *gin.Context) {
	//ph.service.Save(ctx)
}

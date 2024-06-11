package services

import (
	"context"
	iErrors "products/internal/errors"
	"products/internal/models"
)

type ProductService interface {
	Save(ctx context.Context, product models.Product) (*models.Product, *iErrors.ApplicationError)
}

type ProductServiceImpl struct {
}

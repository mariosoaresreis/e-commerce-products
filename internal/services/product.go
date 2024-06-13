package services

import (
	"context"
	iErrors "products/internal/errors"
	"products/internal/models"
	"products/internal/repository"
)

type ProductService interface {
	Save(ctx context.Context, product models.Product) (*models.Product, *iErrors.ApplicationError)
}

type ProductServiceImpl struct {
	repository repository.ProductRepository
}

func (p *ProductServiceImpl) Save(ctx context.Context, product models.Product) (*models.Product,
	*iErrors.ApplicationError) {
	return nil, nil
}

func NewProductService() ProductService {
	return &ProductServiceImpl{}
}

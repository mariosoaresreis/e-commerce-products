package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"products/internal/models"
)

type ProductRepository interface {
	Save(product *models.Product) *models.Product
	Search(searchString string) ([]*models.Product, error)
}

type ProductRepositoryImpl struct {
	client *mongo.Client
}

func (r *ProductRepositoryImpl) Save(product *models.Product) *models.Product {
	return nil
}

func (r *ProductRepositoryImpl) Search(searchString string) ([]*models.Product, error) {
	return nil, nil
}

func NewProductRepository(c *mongo.Client) ProductRepository {
	return &ProductRepositoryImpl{
		client: c,
	}
}

package feature

import (
	"main/domain/product/model"
	"main/domain/product/repository"
)

// ProductFeature is an interface that defines the methods of the product feature.
type ProductFeature interface {
	GetProductList() ([]model.Product, error)
}

type productFeature struct {
	productRepository repository.ProductRepository
}

// NewProductFeature creates a new instance of ProductFeature.
func NewProductFeature(productRepository repository.ProductRepository) ProductFeature {
	return &productFeature{
		productRepository: productRepository,
	}
}


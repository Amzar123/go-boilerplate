package repository

import "main/domain/product/model"

type ProductRepository interface {
	GetProductList() ([]model.Product, error)
}

type productRepository struct {
	products []model.Product
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		products: []model.Product{
			{ID: 1, Name: "Product 1", Price: 100},
			{ID: 2, Name: "Product 2", Price: 200},
			{ID: 3, Name: "Product 3", Price: 300},
		},
	}
}

func (r productRepository) GetProductList() ([]model.Product, error) {
	return r.products, nil
}
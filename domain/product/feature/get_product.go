package feature

import (
	"main/domain/product/model"
)

func (f productFeature) GetProductList() ([]model.Product, error) {
	return f.productRepository.GetProductList()
}
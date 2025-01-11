package http

import (
	"main/delivery/container"
	"main/domain/product"
)

type handler struct {
	productHandler product.ProductHandler
}

func NewHandler(
	container container.Container,
) handler {
	return handler{
		productHandler: product.NewProductHandler(container.ProductFeature),
	}
}
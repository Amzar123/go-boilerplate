package container

import (
	productFeature "main/domain/product/feature"
	productRepo "main/domain/product/repository"

	"fmt"
)

type Container struct {
	ProductFeature           productFeature.ProductFeature
}

func SetupContainer() Container {
	fmt.Println("Starting new container...")

	fmt.Println("Loading repository's...")
	productRepo := productRepo.NewProductRepository()

	fmt.Println("Loading feature's...")
	productFeature := productFeature.NewProductFeature(productRepo)

	return Container{
		ProductFeature: productFeature,
	}
}

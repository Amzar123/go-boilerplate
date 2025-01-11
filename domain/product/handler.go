package product

import (
	"main/domain/product/feature"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	GetProductList(c *fiber.Ctx) error
}

type productHandler struct {
	productFeature feature.ProductFeature
}

func NewProductHandler(productFeature feature.ProductFeature) ProductHandler {
	return &productHandler{
		productFeature: productFeature,
	}
}

func (h productHandler)GetProductList(c *fiber.Ctx) error {
	products, err := h.productFeature.GetProductList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

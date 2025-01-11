package http

import (
	"log"

	_ "main/docs/api"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Approval Management System API
// @version 1.0
// @description This is API Documentation for AMS
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9000
// @BasePath /
func RouteGroup(app *fiber.App, handler handler) {
	log.Println("Setting up routes...")

	v1 := app.Group("/v1")

	// Swagger Documentation Route
	v1.Get("/docs/*", fiberSwagger.WrapHandler)

	// Product Routes
	product := v1.Group("/product")
	product.Get("/list", handler.productHandler.GetProductList)
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropazello/lojinha-product-catalog-service/cmd/web_api/controllers"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/src/infra/databases/dynamo"
)

func main() {
	productRepository := dynamo.ProductRepository{}
	findProductUseCase := usecases.NewFindProductUseCase(&productRepository)
	createProductUseCase := usecases.NewCreateProductUseCase(&productRepository)
	productsController := controllers.NewProductsController(
		findProductUseCase,
		createProductUseCase,
	)

	router := gin.Default()
	router.GET("/product/:id", productsController.Show)
	router.POST("/products", productsController.Create)
	router.Run()
}

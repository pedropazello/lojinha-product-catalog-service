package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropazello/lojinha-product-catalog-service/cmd/web_api/controllers"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/infra/databases/dynamo"
)

func main() {
	productRepository := dynamo.ProductRepository{}
	findProductUseCase := usecases.NewFindProductUseCase(&productRepository)
	createProductUseCase := usecases.NewCreateProductUseCase(&productRepository)
	updateProductUseCase := usecases.NewUpdateProductUseCase(&productRepository)
	deleteProductUseCase := usecases.NewDeleteProductUseCase(&productRepository)

	productsController := controllers.NewProductsController(
		findProductUseCase,
		createProductUseCase,
		updateProductUseCase,
		deleteProductUseCase,
	)

	router := gin.Default()
	router.POST("/products", productsController.Create)
	router.GET("/product/:id", productsController.Show)
	router.PATCH("/product/:id", productsController.Update)
	router.DELETE("/product/:id", productsController.Delete)

	router.Run()
}

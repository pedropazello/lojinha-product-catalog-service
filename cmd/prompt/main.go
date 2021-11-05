package main

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/src/delivery/prompt"
	"github.com/pedropazello/lojinha-product-catalog-service/src/infra/databases/dynamo"
	"github.com/pedropazello/lojinha-product-catalog-service/src/infra/serializers"
	"github.com/pedropazello/lojinha-product-catalog-service/src/infra/stdout"
)

func main() {
	productRepository := &dynamo.ProductRepository{}
	createProductUseCase := usecases.NewCreateProductUseCase(productRepository)
	findProductUseCase := usecases.NewFindProductUseCase(productRepository)
	updateProductUseCase := usecases.NewUpdateProductUseCase(productRepository)
	deleteProductUseCase := usecases.NewDeleteProductUseCase(productRepository)

	prompt := prompt.NewHandler(
		createProductUseCase,
		findProductUseCase,
		updateProductUseCase,
		deleteProductUseCase,
		&stdout.CLI{},
		&serializers.JSONSerializer{},
	)

	prompt.Start()
}

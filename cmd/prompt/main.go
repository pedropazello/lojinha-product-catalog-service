package main

import (
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/delivery/prompt"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/infra/databases/dynamo"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/infra/serializers"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/infra/stdout"
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

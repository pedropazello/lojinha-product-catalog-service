package usecases

import (
	"github.com/google/uuid"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/repositories"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases/interfaces"
)

func NewFindProductUseCase(productRepository repositories.IProductRepository) interfaces.IFindProductUseCase {
	return &FindProductUseCase{
		productRepository: productRepository,
	}
}

type FindProductUseCase struct {
	productRepository repositories.IProductRepository
}

func (r FindProductUseCase) Execute(productId uuid.UUID) (output.ProductData, error) {
	output := output.ProductData{
		ID: uuid.Nil,
	}

	fetchedProduct, err := r.productRepository.GetById(productId)

	if err != nil {
		return output, err
	}

	output.ID = fetchedProduct.ID
	output.Name = fetchedProduct.Name
	output.Description = fetchedProduct.Description

	return output, nil
}

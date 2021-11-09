package usecases

import (
	"github.com/google/uuid"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/output"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/repositories"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases/interfaces"
)

func NewFindProductUseCase(productRepository repositories.IProductRepository) interfaces.IFindProductUseCase {
	return &FindProductUseCase{
		productRepository: productRepository,
	}
}

type FindProductUseCase struct {
	productRepository repositories.IProductRepository
}

func (r FindProductUseCase) Execute(productId string) (output.ProductData, error) {
	output := output.ProductData{
		ID: uuid.Nil.String(),
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

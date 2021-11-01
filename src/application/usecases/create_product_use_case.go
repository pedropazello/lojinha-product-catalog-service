package usecases

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/repositories"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases/interfaces"
	"github.com/pedropazello/lojinha-product-catalog-service/src/domain/entities"
)

type createProductUseCase struct {
	productRepository repositories.IProductRepository
}

func NewCreateProductUseCase(productRepository repositories.IProductRepository) interfaces.ICreateProductUseCase {
	return &createProductUseCase{
		productRepository: productRepository,
	}
}

func (p *createProductUseCase) Execute(productInputData *input.ProductData) (*output.ProductData, error) {
	product := entities.Product{
		Name:        productInputData.Name,
		Description: productInputData.Description,
	}

	createdProduct, err := p.productRepository.Create(&product)

	outputProduct := output.ProductData{
		ID:          createdProduct.ID,
		Name:        createdProduct.Name,
		Description: createdProduct.Description,
	}

	if err != nil {
		return &outputProduct, err
	}

	return &outputProduct, nil
}

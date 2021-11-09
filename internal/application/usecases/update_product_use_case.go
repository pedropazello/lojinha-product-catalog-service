package usecases

import (
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/output"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/repositories"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases/interfaces"
)

func NewUpdateProductUseCase(productRepository repositories.IProductRepository) interfaces.IUpdateProductUseCase {
	return &UpdateProductUseCase{
		productRepository: productRepository,
	}
}

type UpdateProductUseCase struct {
	productRepository repositories.IProductRepository
}

func (u *UpdateProductUseCase) Execute(inputProduct *input.ProductData) (output.ProductData, error) {
	output := output.ProductData{}

	product, err := u.productRepository.GetById(inputProduct.ID)
	if err != nil {
		return output, err
	}

	product.Name = inputProduct.Name
	product.Description = inputProduct.Description

	err = u.productRepository.Save(product)
	if err != nil {
		return output, err
	}

	output.ID = product.ID
	output.Name = product.Name
	output.Description = product.Description

	return output, nil
}

package usecases

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/repositories"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases/interfaces"
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

	if inputProduct.Name != "" {
		product.Name = inputProduct.Name
	}

	if inputProduct.Description != "" {
		product.Description = inputProduct.Description
	}

	err = u.productRepository.Save(product)
	if err != nil {
		return output, err
	}

	output.ID = product.ID
	output.Name = product.Name
	output.Description = product.Description

	return output, nil
}

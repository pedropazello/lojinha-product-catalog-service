package interfaces

import (
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/output"
)

type ICreateProductUseCase interface {
	Execute(product *input.ProductData) (*output.ProductData, error)
}

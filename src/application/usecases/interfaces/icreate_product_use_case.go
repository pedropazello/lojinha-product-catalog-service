package interfaces

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
)

type ICreateProductUseCase interface {
	Execute(product *input.ProductData) (*output.ProductData, error)
}

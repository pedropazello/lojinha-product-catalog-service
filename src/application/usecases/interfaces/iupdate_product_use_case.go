package interfaces

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/input"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
)

type IUpdateProductUseCase interface {
	Execute(inputProduct *input.ProductData) (output.ProductData, error)
}

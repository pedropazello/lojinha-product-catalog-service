package interfaces

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
)

type IFindProductUseCase interface {
	Execute(productId string) (output.ProductData, error)
}

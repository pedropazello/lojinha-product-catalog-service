package interfaces

import (
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/output"
)

type IFindProductUseCase interface {
	Execute(productId string) (output.ProductData, error)
}

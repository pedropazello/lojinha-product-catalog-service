package interfaces

import (
	"github.com/google/uuid"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/ports/output"
)

type IFindProductUseCase interface {
	Execute(productId uuid.UUID) (output.ProductData, error)
}

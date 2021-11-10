package interfaces

import "github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/input"

type IImportProductUseCase interface {
	Execute([]input.ProductData) (int, []error)
}

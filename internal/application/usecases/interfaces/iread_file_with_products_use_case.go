package interfaces

import "github.com/pedropazello/lojinha-product-catalog-service/internal/application/ports/output"

type IReadFileWithProductsUseCase interface {
	Execute(file []byte) ([]output.ProductData, error)
}

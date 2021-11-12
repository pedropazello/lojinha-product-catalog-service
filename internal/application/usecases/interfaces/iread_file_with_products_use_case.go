package interfaces

import (
	"github.com/pedropazello/lojinha-product-catalog-service/internal/domain/entities"
)

type IReadFileWithProductsUseCase interface {
	Execute(file []byte) ([]entities.Product, error)
}

package interfaces

import (
	"github.com/pedropazello/lojinha-product-catalog-service/internal/domain/entities"
)

type IImportProductsUseCase interface {
	Execute([]entities.Product) (int, []error)
}

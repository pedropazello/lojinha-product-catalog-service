package repositories

import (
	"github.com/google/uuid"
	"github.com/pedropazello/lojinha-product-catalog-service/src/domain/entities"
)

type IProductRepository interface {
	Create(product *entities.Product) (*entities.Product, error)
	GetById(id uuid.UUID) (*entities.Product, error)
	Save(product *entities.Product) error
	Delete(id uuid.UUID) error
}

package repositories

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/domain/entities"
)

type IProductRepository interface {
	Create(product *entities.Product) (*entities.Product, error)
	GetById(id string) (*entities.Product, error)
	Save(product *entities.Product) error
	Delete(id string) error
}

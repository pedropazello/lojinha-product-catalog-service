package parsers

import "github.com/pedropazello/lojinha-product-catalog-service/internal/domain/entities"

type IProductsFileParser interface {
	Parse(file []byte) ([]entities.Product, error)
}

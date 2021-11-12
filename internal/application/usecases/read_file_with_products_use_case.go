package usecases

import (
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/parsers"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/domain/entities"
)

type readFileWithProductsUseCase struct {
	parser parsers.IProductsFileParser
}

func NewReadFileWithProductsUseCase(parser parsers.IProductsFileParser) *readFileWithProductsUseCase {
	return &readFileWithProductsUseCase{
		parser: parser,
	}
}

func (r *readFileWithProductsUseCase) Execute(file []byte) ([]entities.Product, error) {
	products := []entities.Product{}

	parsedProducts, err := r.parser.Parse(file)
	if err != nil {
		return products, err
	}

	for _, product := range parsedProducts {
		products = append(products, entities.Product{
			Name:        product.Name,
			Description: product.Description,
		})
	}

	return products, nil
}

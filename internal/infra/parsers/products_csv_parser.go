package parsers

import (
	"bytes"
	"encoding/csv"

	"github.com/pedropazello/lojinha-product-catalog-service/internal/domain/entities"
)

type ProductsCSVParser struct {
}

func (parser *ProductsCSVParser) Parse(file []byte) ([]entities.Product, error) {
	products := []entities.Product{}
	reader := bytes.NewReader(file)

	csvLines, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return products, err
	}

	for _, line := range csvLines[1:] {
		product := entities.Product{
			Name:        line[0],
			Description: line[1],
		}

		products = append(products, product)
	}

	return products, nil
}

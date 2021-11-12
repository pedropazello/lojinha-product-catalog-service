package usecases

import (
	"fmt"

	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/repositories"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases/interfaces"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/domain/entities"
)

func NewImportProductsUseCase(repository repositories.IProductRepository) interfaces.IImportProductsUseCase {
	return &importProductsUseCase{
		repository: repository,
	}
}

type importProductsUseCase struct {
	repository repositories.IProductRepository
}

func (u *importProductsUseCase) Execute(productsInput []entities.Product) (int, []error) {
	productsInported := 0
	errs := []error{}

	for _, productInput := range productsInput {
		product := entities.Product{
			Name:        productInput.Name,
			Description: productInput.Description,
		}

		err := u.repository.Save(&product)

		if err != nil {
			err = fmt.Errorf("%w; Product failed:"+productInput.Name, err)
			errs = append(errs, err)
		} else {
			productsInported += 1
		}
	}

	return productsInported, errs
}

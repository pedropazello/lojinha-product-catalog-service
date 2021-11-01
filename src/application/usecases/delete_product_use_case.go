package usecases

import (
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/repositories"
	"github.com/pedropazello/lojinha-product-catalog-service/src/application/usecases/interfaces"
)

func NewDeleteProductUseCase(repo repositories.IProductRepository) interfaces.IDeleteProductUseCase {
	return &DeleteProductUseCase{
		repository: repo,
	}
}

type DeleteProductUseCase struct {
	repository repositories.IProductRepository
}

func (d *DeleteProductUseCase) Execute(id string) error {
	err := d.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

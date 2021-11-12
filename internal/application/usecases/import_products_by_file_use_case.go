package usecases

import (
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/storages"
	"github.com/pedropazello/lojinha-product-catalog-service/internal/application/usecases/interfaces"
)

type importProductsByFileUseCase struct {
	fileStorage                 storages.IFileStorage
	readFileWithProductsUseCase interfaces.IReadFileWithProductsUseCase
	importProductsUseCase       interfaces.IImportProductsUseCase
}

func NewImportProductsByFileUseCase(fileStorage storages.IFileStorage,
	readFileWithProductsUseCase interfaces.IReadFileWithProductsUseCase,
	importProductsUseCase interfaces.IImportProductsUseCase) interfaces.IImportProductsByFileUseCase {

	return &importProductsByFileUseCase{
		fileStorage:                 fileStorage,
		readFileWithProductsUseCase: readFileWithProductsUseCase,
		importProductsUseCase:       importProductsUseCase,
	}
}

func (i *importProductsByFileUseCase) Execute(fileName string) error {
	file, err := i.fileStorage.GetFile(fileName)
	if err != nil {
		return err
	}

	products, err := i.readFileWithProductsUseCase.Execute(file)
	if err != nil {
		return err
	}

	i.importProductsUseCase.Execute(products)
	return nil
}

package interfaces

type IImportProductsByFileUseCase interface {
	Execute(fileName string) error
}

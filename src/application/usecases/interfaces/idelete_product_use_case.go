package interfaces

type IDeleteProductUseCase interface {
	Execute(id string) error
}

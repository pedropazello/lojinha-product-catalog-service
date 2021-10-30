package interfaces

import "github.com/google/uuid"

type IDeleteProductUseCase interface {
	Execute(id uuid.UUID) error
}
